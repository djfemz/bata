package utils

import "strconv"

const (
	PAYSTACK_TRANSACTION_URL     = "https://api.paystack.co/transaction/initialize"
	AUTHORIZATION_HEADER         = "Authorization"
	CONTENT_TYPE                 = "Content-Type"
	SIMPLE_BANK_EMAIL_ADDRESS    = "simplebank@email.com"
	APPLICATION_JSON_VALUE       = "application/json"
	CREDIT_TRANSACTION           = "CREDIT"
	DEBIT_TRANSACTION            = "DEBIT"
	ACCOUNT_CREATED_SUCCESSFULLY = "account created successfully"
)

func ConvertAmountToString(amount float64) string {
	return strconv.FormatFloat(amount, 'g', 2, 64)
}
