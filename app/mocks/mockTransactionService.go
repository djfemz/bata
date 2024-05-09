package mocks

import (
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/dtos/responses"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	mock.Mock
}

func (transactionService *MockTransactionService) CreateTransaction(request *requests.PaystackTransactionRequest) (*responses.PaystackTransactionResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(request)
	args := transactionService.Called(request)
	if err != nil {
		err = args.Error(1)
		return nil, err
	}

	return &responses.PaystackTransactionResponse{
		Status:  true,
		Message: "success",
		Data: responses.Data{
			AuthorizationUrl: "www.paystack.com",
			AccessCode:       "12345",
			Reference:        "12345",
		},
	}, nil
}
