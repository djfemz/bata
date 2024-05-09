package services

import (
	"encoding/json"
	"github.com/djfemz/simple_bank/app/appErrors"
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/dtos/responses"
	"github.com/djfemz/simple_bank/app/models"
	"github.com/djfemz/simple_bank/app/repositories"
	"github.com/go-playground/validator/v10"
	"gopkg.in/jeevatkm/go-model.v1"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type TransactionService interface {
	PerformTransaction(debitTransactionRequest *requests.CreateTransactionRequest) (*responses.TransactionResponse, error)
}

type bankTransactionService struct {
	transactionRepository      repositories.TransactionRepository
	accountService             AccountService
	paystackTransactionService PaystackTransactionService
}

func NewTransactionService(paystackService PaystackTransactionService) TransactionService {
	return &bankTransactionService{
		repositories.NewTransactionRepository(),
		NewAccountService(),
		paystackService,
	}
}

func (transactionService *bankTransactionService) PerformTransaction(debitTransactionRequest *requests.CreateTransactionRequest) (*responses.TransactionResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())

	var savedTransaction *models.Transaction
	var err error
	err = validate.Struct(debitTransactionRequest)
	if err != nil {
		return nil, err
	}
	transaction := &models.Transaction{}
	errs := model.Copy(transaction, debitTransactionRequest)
	if len(errs) != 0 {
		return nil, appErrors.NewTransactionFailedError()
	}
	transactionRepository := transactionService.transactionRepository
	databaseConnection, err := transactionRepository.GetDatabaseConnection()
	if err != nil {
		return nil, appErrors.NewCreateTransactionFailedError()
	}
	err = databaseConnection.Transaction(func(tx *gorm.DB) error {
		savedTransaction, err = transactionRepository.Save(transaction)
		if err != nil {
			return err
		}
		paystackRequest := createPaystackTransactionRequest(savedTransaction)
		response, err := transactionService.paystackTransactionService.CreateTransaction(paystackRequest)
		if err != nil {
			return err
		}
		savedTransaction.Reference = response.Data.Reference
		accountService := transactionService.accountService
		err = accountService.UpdateBalanceWith(savedTransaction)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	log.Println("transaction successful: ", debitTransactionRequest)
	return buildTransactionResponse(savedTransaction), nil
}

func buildTransactionResponse(transaction *models.Transaction) *responses.TransactionResponse {
	transactionResponse := &responses.TransactionResponse{}
	transactionResponse.AccountNumber = transaction.AccountNumber
	transactionResponse.Reference = transaction.Reference
	transactionResponse.Amount = transaction.Amount
	return transactionResponse
}

func extractTransactionResponseFrom(res *http.Response) (*responses.PaystackTransactionResponse, error) {
	paystackTransactionResponse := &responses.PaystackTransactionResponse{}
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(paystackTransactionResponse)
	if err != nil {
		return nil, err
	}

	return paystackTransactionResponse, nil
}
