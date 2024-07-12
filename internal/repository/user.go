package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/prastamaha/envelope-encryption-example/internal/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(ctx context.Context, user *model.User) (string, error) {
	query := `INSERT INTO users (username, name, gender, phone, address, consented) VALUES ($1, $2, $3, $4, $5, $6)`
	stmt, err := r.db.PreparexContext(ctx, query)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	_, err = stmt.ExecContext(ctx, user.Username, user.Name, user.Gender, user.Phone, user.Address, user.Consented)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return user.Username, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `SELECT username, name, gender, phone, address, consented, created_at FROM users WHERE username = $1`
	stmt, err := r.db.PreparexContext(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	row := stmt.QueryRowxContext(ctx, username)

	var user model.User
	if err := row.StructScan(&user); err != nil {
		fmt.Println(err)
		return nil, err
	}

	user.Name = string(user.Name)
	user.Gender = string(user.Gender)
	user.Phone = string(user.Phone)
	user.Address = string(user.Address)

	return &user, nil
}
