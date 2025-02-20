package domain

type User struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}
