package requests

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateTransactionRequest struct {
	AccountNumber string  `json:"account_id"`
	Type          string  `json:"transaction_type" validate:"oneof=DEBIT CREDIT"`
	Amount        float64 `json:"amount" validate:"gte=1"`
}

type PaystackTransactionRequest struct {
	Email  string `json:"email" validate:"required"`
	Amount string `json:"amount" validate:"required"`
}
