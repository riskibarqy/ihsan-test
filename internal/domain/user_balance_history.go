package domain

type UserBalanceHistory struct {
	ID              int    `json:"id,omitempty"`
	UserID          int    `json:"user_id"`
	NoRekening      string `json:"no_rekening"`
	PreviousBalance int    `json:"previous_balance"`
	ChangeAmount    int    `json:"change_amount"`
	NewBalance      int    `json:"new_balance"`
	TransactionType string `json:"transaction_type"`
	TransactionID   string `json:"transaction_id"`
	CreatedAt       int    `json:"created_at"`
}
