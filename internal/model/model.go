package model

type User struct {
	Username  string `db:"username"`
	Name      string `db:"encrypted_name"`
	Gender    string `db:"encrypted_gender"`
	Phone     string `db:"encrypted_phone"`
	Address   string `db:"encrypted_address"`
	DEK       []byte `db:"encrypted_dek"`
	Consented bool   `db:"consented"`
	CreatedAt string `db:"created_at"`
}

type DEK struct {
	Ciphertext string `json:"ciphertext"`
	Plaintext  string `json:"plaintext"`
}
