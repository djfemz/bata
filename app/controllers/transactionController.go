package controllers

import (
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/dtos/responses"
	"github.com/djfemz/simple_bank/app/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TransactionController struct {
	transactionService services.TransactionService
}

func NewTransactionController() *TransactionController {
	return &TransactionController{
		services.NewTransactionService(services.NewAppPayStackTransactionService()),
	}
}

func (transactionController *TransactionController) PerformTransaction(ctx *gin.Context) {
	createTransactionRequest := &requests.CreateTransactionRequest{}
	err := ctx.BindJSON(createTransactionRequest)
	log.Println("transaction request: ", createTransactionRequest)
	if err != nil {
		log.Println("transaction  failed: ", err)
		ctx.JSON(http.StatusBadRequest, &responses.Response[string]{Data: err.Error()})
		return
	}
	performTransactionResponse, err := transactionController.transactionService.PerformTransaction(createTransactionRequest)
	if err != nil {
		log.Println("transaction failed: ", err)
		ctx.JSON(http.StatusBadRequest, &responses.Response[string]{Data: err.Error()})
		return
	}
	log.Println("transaction successful: ", createTransactionRequest)
	ctx.JSON(http.StatusCreated, responses.Response[responses.TransactionResponse]{Data: *performTransactionResponse, Success: true})
}
