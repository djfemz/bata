package appErrors

import "errors"

const (
	ERR_TRANSACTION_NOT_FOUND                   = "transaction not found"
	ERR_ACCOUNT_NOT_FOUND                       = "account not found"
	ERR_TRANSACTION_FAILED                      = "transaction failed, try again later"
	ERR_TRANSACTION_FAILED_INSUFFICIENT_BALANCE = "transaction failed due to insufficient funds"
	ERR_PERFORMING_TRANSACTION                  = "failed to perform transaction. Try again later"
	ERR_ACCOUNT_CREATION_FAILED                 = "account creation failed"
)

func NewTransactionNotFoundError() error {
	return errors.New(ERR_TRANSACTION_NOT_FOUND)
}

func NewCreateTransactionFailedError() error {
	return errors.New(ERR_PERFORMING_TRANSACTION)
}

func NewAccountCreationFailedError() error {
	return errors.New(ERR_ACCOUNT_CREATION_FAILED)
}

func NewTransactionFailedError() error {
	return errors.New(ERR_TRANSACTION_FAILED)
}

func NewTransactionFailedInsufficientFundsError() error {
	return errors.New(ERR_TRANSACTION_FAILED_INSUFFICIENT_BALANCE)
}
