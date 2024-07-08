package model

import (
	"context"
)

type EnvelopeEncryption interface {
	GenerateDEK() (*DEK, error)
	DecryptDEK(ciphertext []byte) ([]byte, error)
	Encrypt(dek []byte, plaintext []byte) ([]byte, error)
	Decrypt(dek []byte, ciphertext []byte) ([]byte, error)
}

type UserRepository interface {
	RegisterUser(ctx context.Context, user *User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}
