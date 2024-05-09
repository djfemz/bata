package repositories

import (
	"github.com/djfemz/simple_bank/app/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction *models.Transaction) (*models.Transaction, error)
	GetDatabaseConnection() (*gorm.DB, error)
}

type bankTransactionRepository struct{}

func NewTransactionRepository() TransactionRepository {
	return &bankTransactionRepository{}
}

func (transactionRepository *bankTransactionRepository) Save(transaction *models.Transaction) (*models.Transaction, error) {
	if databaseConnection, err = Connect(); err != nil {
		return nil, err
	}
	err = databaseConnection.Transaction(func(tx *gorm.DB) error {
		if err = tx.Save(transaction).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return transaction, err
}

func (transactionRepository *bankTransactionRepository) GetDatabaseConnection() (*gorm.DB, error) {
	if databaseConnection, err = Connect(); err != nil {
		return nil, err
	}
	return databaseConnection, nil
}
