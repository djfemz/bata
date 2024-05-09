package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/djfemz/simple_bank/app/controllers"
	"github.com/djfemz/simple_bank/app/dtos/requests"
	"github.com/djfemz/simple_bank/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var transactionController = controllers.NewTransactionController()

func TestCreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		router := gin.Default()
		router.POST("/api/v1/transaction", transactionController.PerformTransaction)
		writer := httptest.NewRecorder()
		createAccountRequest := &requests.CreateTransactionRequest{
			AccountNumber: "2212160567",
			Type:          utils.CREDIT_TRANSACTION,
			Amount:        1000.00,
		}
		data, _ := json.Marshal(createAccountRequest)
		request, _ := http.NewRequest(http.MethodPost, "/api/v1/transaction", bytes.NewReader(data))
		request.Header.Add("Content-Type", "application/json")
		router.ServeHTTP(writer, request)
		log.Println(writer.Body.String())
		assert.Equal(t, http.StatusCreated, writer.Code)
		assert.NotNil(t, writer.Body)
	})

}
