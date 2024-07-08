package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/prastamaha/envelope-encryption-example/internal/model"
)

type UserRepository struct {
	db                 *sqlx.DB
	envelopeEncryption model.EnvelopeEncryption
}

func NewUserRepository(db *sqlx.DB, ee model.EnvelopeEncryption) *UserRepository {
	return &UserRepository{
		db:                 db,
		envelopeEncryption: ee,
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user *model.User) (string, error) {
	query := `INSERT INTO users (username, encrypted_name, encrypted_gender, encrypted_phone, encrypted_address, encrypted_dek, consented) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	stmt, err := r.db.PreparexContext(ctx, query)
	if err != nil {
		return "", err
	}

	dek, err := r.envelopeEncryption.GenerateDEK()
	if err != nil {
		return "", err
	}

	encryptedName, err := r.envelopeEncryption.Encrypt([]byte(dek.Plaintext), []byte(user.Name))
	if err != nil {
		return "", err
	}

	encryptedGender, err := r.envelopeEncryption.Encrypt([]byte(dek.Plaintext), []byte(user.Gender))
	if err != nil {
		return "", err
	}

	encryptedPhone, err := r.envelopeEncryption.Encrypt([]byte(dek.Plaintext), []byte(user.Phone))
	if err != nil {
		return "", err
	}

	encryptedAddress, err := r.envelopeEncryption.Encrypt([]byte(dek.Plaintext), []byte(user.Address))
	if err != nil {
		return "", err
	}

	_, err = stmt.ExecContext(ctx, user.Username, encryptedName, encryptedGender, encryptedPhone, encryptedAddress, dek.Ciphertext, user.Consented)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `SELECT username, encrypted_name, encrypted_gender, encrypted_phone, encrypted_address, encrypted_dek, consented, created_at FROM users WHERE username = $1`
	stmt, err := r.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowxContext(ctx, username)

	var user model.User
	if err := row.StructScan(&user); err != nil {
		return nil, err
	}

	plainDek, err := r.envelopeEncryption.DecryptDEK(user.DEK)
	if err != nil {
		return nil, err
	}

	plainName, err := r.envelopeEncryption.Decrypt(plainDek, []byte(user.Name))
	if err != nil {
		return nil, err
	}

	plainGender, err := r.envelopeEncryption.Decrypt(plainDek, []byte(user.Gender))
	if err != nil {
		return nil, err
	}

	plainPhone, err := r.envelopeEncryption.Decrypt(plainDek, []byte(user.Phone))
	if err != nil {
		return nil, err
	}

	plainAddress, err := r.envelopeEncryption.Decrypt(plainDek, []byte(user.Address))
	if err != nil {
		return nil, err
	}

	user.Name = string(plainName)
	user.Gender = string(plainGender)
	user.Phone = string(plainPhone)
	user.Address = string(plainAddress)

	return &user, nil
}
