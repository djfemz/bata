package test

import (
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/services"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var accountService services.AccountService

func TestCreateAccount(t *testing.T) {
	accountService = services.NewAccountService()
	createAccountRequest := &requests.CreateAccountRequest{
		Username: "john@email.com",
		Password: "12345",
	}
	response, err := accountService.CreateAccount(createAccountRequest)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Contains(t, response.Message, "success")
	assert.NotNil(t, response.AccountNumber)
}

func TestGetAccountById(t *testing.T) {
	accountService = services.NewAccountService()
	account, err := accountService.GetAccountBy("2212160567")
	log.Println(account)
	assert.NotNil(t, account)
	assert.Nil(t, err)
}
