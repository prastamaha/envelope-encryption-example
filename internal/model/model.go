package model

type User struct {
	Username  string `db:"username"`
	Name      string `db:"name"`
	Gender    string `db:"gender"`
	Phone     string `db:"phone"`
	Address   string `db:"address"`
	Consented bool   `db:"consented"`
	CreatedAt string `db:"created_at"`
}

type DEK struct {
	Ciphertext string `json:"ciphertext"`
	Plaintext  string `json:"plaintext"`
}
