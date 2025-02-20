package datatransfers

type UserGetBalanceResponse struct {
	NoRekening string `json:"no_rekening"`
	Saldo      int    `json:"saldo"`
}
