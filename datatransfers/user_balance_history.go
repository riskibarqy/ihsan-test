package datatransfers

import "github.com/riskibarqy/ihsan-test/internal/constants"

type UserBalanceHistoryTXResponse struct {
	NoRekening string `json:"no_rekening"`
	Saldo      int    `json:"saldo"`
}

type UserBalanceHistoryTXRequest struct {
	NoRekening      string                    `json:"no_rekening"`
	Nominal         int                       `json:"nominal"`
	TransactionType constants.TransactionType `json:"-"`
}
