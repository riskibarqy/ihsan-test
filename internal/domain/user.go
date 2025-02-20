package domain

type User struct {
	ID         int    `json:"id,omitempty"`
	Nama       string `json:"nama"`
	Nik        string `json:"nik"`
	NoHp       string `json:"no_hp"`
	NoRekening string `json:"no_rekening"`
	Balance    int    `json:"balance,omitempty"`
	CreatedAt  int    `json:"created_at,omitempty"`
	UpdatedAt  int    `json:"updated_at,omitempty"`
	DeletedAt  int    `json:"deleted_at,omitempty"`
}
