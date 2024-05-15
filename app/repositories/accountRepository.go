package repositories

import (
	"fmt"
	"github.com/djfemz/simple_bank/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var databaseConnection *gorm.DB
var err error

type AccountRepository interface {
	Save(account *models.Account) (*models.Account, error)
	FindByAccountNumber(number string) (*models.Account, error)
}

type bankAccountRepository struct{}

func NewAccountRepository() AccountRepository {
	return &bankAccountRepository{}
}

func (accountRepository *bankAccountRepository) Save(account *models.Account) (*models.Account, error) {
	databaseConnection, err = Connect()
	if err != nil {
		return nil, err
	}
	return account, databaseConnection.Save(account).Error
}

func (accountRepository *bankAccountRepository) FindByAccountNumber(accountNumber string) (*models.Account, error) {
	log.Println("account nuumber", accountNumber)
	databaseConnection, err = Connect()
	account := &models.Account{}
	err = databaseConnection.Where("account_number=?", accountNumber).First(account).Error
	if err != nil {
		return nil, err
	}
	return account, err
}

func Connect() (*gorm.DB, error) {
	dsn, err := buildDbConnectionString()
	if err != nil {
		return nil, err
	}
	databaseConnection, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = databaseConnection.AutoMigrate(&models.Account{}, &models.Transaction{})
	if err != nil {
		return nil, err
	}
	return databaseConnection, nil
}

func buildDbConnectionString() (string, error) {
	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=Africa/Lagos sslmode=disable", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"), port), nil
}
