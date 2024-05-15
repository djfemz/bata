package main

import (
	"github.com/djfemz/simple_bank/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading environment variables: ", err)
	}
	router := gin.Default()
	transactionController := controllers.NewTransactionController()
	router.POST("/api/v1/transaction", transactionController.PerformTransaction)
	err = router.Run(":8000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
