package responses

type Response[T any] struct {
	Data    T
	Success bool `json:"success"`
}

type CreateAccountResponse struct {
	Message       string `json:"message"`
	AccountNumber string `json:"account_id"`
}

type AccountResponse struct {
	AccountNumber string  `json:"account_id"`
	Balance       float64 `json:"account_balance"`
}

type Data struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type TransactionResponse struct {
	AccountNumber string  `json:"account_id"`
	Reference     string  `json:"reference"`
	Amount        float64 `json:"amount"`
}

type PaystackTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data
}
