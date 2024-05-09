package services

import (
	"bytes"
	"encoding/json"
	"github.com/djfemz/simple_bank/app/appErrors"
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/dtos/responses"
	"github.com/djfemz/simple_bank/app/models"
	"github.com/djfemz/simple_bank/app/utils"
	"net/http"
	"os"
)

type PaystackTransactionService interface {
	CreateTransaction(paysatckTransactionRequest *requests.PaystackTransactionRequest) (*responses.PaystackTransactionResponse, error)
}

type AppPayStackTransactionService struct {
}

func NewAppPayStackTransactionService() PaystackTransactionService {
	return &AppPayStackTransactionService{}
}

func (paystackTransactionService *AppPayStackTransactionService) CreateTransaction(paysatckTransactionRequest *requests.PaystackTransactionRequest) (*responses.PaystackTransactionResponse, error) {
	jsonData, _ := json.Marshal(paysatckTransactionRequest)
	request, err := http.NewRequest(http.MethodPost, utils.PAYSTACK_TRANSACTION_URL, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	addHeadersTo(request)

	client := &http.Client{}
	if res, err := client.Do(request); err != nil {
		return nil, err
	} else if res.StatusCode == http.StatusOK {
		transactionResponse, _ := extractTransactionResponseFrom(res)
		return transactionResponse, nil
	} else {
		return nil, appErrors.NewTransactionFailedError()
	}
}

func addHeadersTo(request *http.Request) {
	request.Header.Add(utils.AUTHORIZATION_HEADER, os.Getenv("PAYSTACK_API_KEY"))
	request.Header.Add(utils.CONTENT_TYPE, utils.APPLICATION_JSON_VALUE)
}

func createPaystackTransactionRequest(transaction *models.Transaction) *requests.PaystackTransactionRequest {
	return &requests.PaystackTransactionRequest{
		Email:  utils.SIMPLE_BANK_EMAIL_ADDRESS,
		Amount: utils.ConvertAmountToString(transaction.Amount),
	}
}
