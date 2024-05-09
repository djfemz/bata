package test

import (
	"github.com/djfemz/simple_bank/app/appErrors"
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/dtos/responses"
	"github.com/djfemz/simple_bank/app/mocks"
	"github.com/djfemz/simple_bank/app/services"
	"github.com/djfemz/simple_bank/app/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

var transactionService services.TransactionService
var mockTransactionService = new(mocks.MockTransactionService)

func TestCreateDebitTransaction(t *testing.T) {
	mockPaystackResponse := &responses.PaystackTransactionResponse{
		Status:  true,
		Message: "success",
		Data: responses.Data{
			Reference:        "12345",
			AuthorizationUrl: "www.paystack.com",
			AccessCode:       "1234",
		},
	}
	accountService = services.NewAccountService()

	mockTransactionService.On("CreateTransaction", mock.AnythingOfType("*requests.PaystackTransactionRequest")).Return(mockPaystackResponse)
	transactionService := services.NewTransactionService(mockTransactionService)
	account, _ := accountService.GetAccountBy("2212160567")
	initialBalance := account.Balance
	debitTransactionRequest := &requests.CreateTransactionRequest{
		AccountNumber: account.AccountNumber,
		Type:          utils.DEBIT_TRANSACTION,
		Amount:        1000.00,
	}
	mockTransactionService.On("CreateTransaction", debitTransactionRequest).Return()
	createTransactionResponse, err := transactionService.PerformTransaction(debitTransactionRequest)
	account, _ = accountService.GetAccountBy(account.AccountNumber)
	assert.Equal(t, initialBalance-debitTransactionRequest.Amount, account.Balance)
	assert.NotNil(t, createTransactionResponse)
	assert.NotEmpty(t, createTransactionResponse.Reference)
	assert.Nil(t, err)
}

func TestThatTransactionWithInvalidAmountFails(t *testing.T) {
	accountService = services.NewAccountService()
	transactionService = services.NewTransactionService(mockTransactionService)
	account, _ := accountService.GetAccountBy("2212160567")
	transactionRequest := &requests.CreateTransactionRequest{
		AccountNumber: account.AccountNumber,
		Type:          utils.DEBIT_TRANSACTION,
		Amount:        20000000.00,
	}
	createTransactionResponse, err := transactionService.PerformTransaction(transactionRequest)
	assert.Nil(t, createTransactionResponse)
	assert.Error(t, err)
}

func TestThatTransactionWithInvalidAccountNumberFails(t *testing.T) {
	accountService = services.NewAccountService()
	transactionService = services.NewTransactionService(mockTransactionService)
	transactionRequest := &requests.CreateTransactionRequest{
		AccountNumber: "22121605",
		Type:          utils.DEBIT_TRANSACTION,
		Amount:        1000.00,
	}
	createTransactionResponse, err := transactionService.PerformTransaction(transactionRequest)
	assert.Nil(t, createTransactionResponse)
	assert.Error(t, err)
}

func TestCreateCreditTransaction(t *testing.T) {
	accountService = services.NewAccountService()
	transactionService = services.NewTransactionService(mockTransactionService)
	account, _ := accountService.GetAccountBy("2212160567")
	initialBalance := account.Balance
	transactionRequest := &requests.CreateTransactionRequest{
		AccountNumber: "2212160567",
		Type:          utils.CREDIT_TRANSACTION,
		Amount:        1000.00,
	}
	createTransactionResponse, err := transactionService.PerformTransaction(transactionRequest)
	account, err = accountService.GetAccountBy("2212160567")
	assert.Equal(t, initialBalance+transactionRequest.Amount, account.Balance)
	assert.NotNil(t, createTransactionResponse)
	assert.NotEmpty(t, createTransactionResponse.Reference)
	assert.Nil(t, err)
}

func TestCreateTransactionWithInvalidTransactionTypeFails(t *testing.T) {
	mockTransactionService.On("CreateTransaction", "*requests.PaystackTransactionRequest").Return(nil, appErrors.NewTransactionFailedError())
	accountService = services.NewAccountService()
	transactionService = services.NewTransactionService(mockTransactionService)
	transactionRequest := &requests.CreateTransactionRequest{
		AccountNumber: "2212160567",
		Type:          "DEBT",
		Amount:        1000.00,
	}
	createTransactionResponse, err := transactionService.PerformTransaction(transactionRequest)
	assert.Nil(t, createTransactionResponse)
	assert.Error(t, err)
}

func TestPerform_Transaction_With_Amount_Less_Than_One_Fails(t *testing.T) {
	transactionRequest := &requests.CreateTransactionRequest{
		AccountNumber: "2212160567",
		Type:          utils.DEBIT_TRANSACTION,
		Amount:        -1000.00,
	}
	accountService = services.NewAccountService()
	mockTransactionService.On("CreateTransaction", mock.AnythingOfType("*requests.PaystackTransactionRequest")).Return(nil, appErrors.NewTransactionFailedError())
	transactionService = services.NewTransactionService(mockTransactionService)

	createTransactionResponse, err := transactionService.PerformTransaction(transactionRequest)
	log.Println(err)
	assert.Nil(t, createTransactionResponse)
	assert.Error(t, err)
}
