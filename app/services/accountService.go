package services

import (
	"errors"
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/dtos/responses"
	"github.com/djfemz/simple_bank/app/models"
	"github.com/djfemz/simple_bank/app/repositories"
	"gopkg.in/jeevatkm/go-model.v1"
	"math/rand"
	"strconv"
	"strings"
)

type AccountService interface {
	CreateAccount(createAccountRequest *requests.CreateAccountRequest) (*responses.CreateAccountResponse, error)
	GetAccountBy(accountNumber string) (*responses.AccountResponse, error)
	UpdateBalanceWith(transaction *models.Transaction) error
}

type bankAccountService struct {
	accountRepository repositories.AccountRepository
}

func NewAccountService() AccountService {
	return &bankAccountService{
		accountRepository: repositories.NewAccountRepository(),
	}
}

func (accountService *bankAccountService) CreateAccount(createAccountRequest *requests.CreateAccountRequest) (*responses.CreateAccountResponse, error) {
	accountRepository := accountService.accountRepository
	account, err := createAccountFor(createAccountRequest)
	if err != nil {
		return nil, err
	}
	account.AccountNumber = generateAccountNumber()
	savedAccount, err := accountRepository.Save(account)
	if err != nil {
		return nil, errors.New("account creation failed")
	}
	return buildCreateAccountResponse(savedAccount)
}

func (accountService *bankAccountService) GetAccountBy(accountNumber string) (*responses.AccountResponse, error) {
	accountRepository := accountService.accountRepository
	foundAccount, err := accountRepository.FindByAccountNumber(accountNumber)
	if err != nil {
		return nil, err
	}

	return buildAccountResponse(foundAccount), nil
}

func (accountService *bankAccountService) UpdateBalanceWith(transaction *models.Transaction) error {
	accountRepository := accountService.accountRepository
	account, err := accountRepository.FindByAccountNumber(transaction.AccountNumber)
	if err != nil {
		return err
	}
	isAccountWithInsufficientFunds := account.Balance < transaction.Amount
	if isAccountWithInsufficientFunds {
		return errors.New("insufficient funds. top-up and try again")
	}
	isDebitTransactionRequest := transaction.Type == "DEBIT"
	isCreditTransactionRequest := transaction.Type == "CREDIT"
	if isDebitTransactionRequest {
		account.Balance = account.Balance - transaction.Amount
	} else if isCreditTransactionRequest {
		account.Balance = account.Balance + transaction.Amount
	}
	_, err = accountRepository.Save(account)
	if err != nil {
		return err
	}
	return nil
}

func buildAccountResponse(account *models.Account) *responses.AccountResponse {
	accountResponse := &responses.AccountResponse{}
	model.Copy(accountResponse, account)

	return accountResponse
}

func createAccountFor(createAccountRequest *requests.CreateAccountRequest) (*models.Account, error) {
	account := &models.Account{}
	errs := model.Copy(account, createAccountRequest)
	isCopyFailed := len(errs) != 0
	if isCopyFailed {
		return nil, errors.New("account creation failed")
	}
	return account, nil
}

func generateAccountNumber() string {
	builder := strings.Builder{}
	for count := 0; count < 10; count++ {
		builder.WriteString(strconv.Itoa(rand.Intn(9)))
	}
	return builder.String()
}

func buildCreateAccountResponse(account *models.Account) (*responses.CreateAccountResponse, error) {
	if account != nil {
		return &responses.CreateAccountResponse{
			Message:       "account creation successful",
			AccountNumber: account.AccountNumber}, nil
	}
	return nil, errors.New("account creation failed")
}
