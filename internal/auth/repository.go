package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateUserTx(
	ctx context.Context,
	tx pgx.Tx,
	name, email, passwordHash, pinHash string,
) (int64, error) {
	var userID int64

	query := `INSERT INTO users(name, email, password_hash, pin_hash) VALUES($1,$2,$3,$4) RETURNING id`

	err := tx.QueryRow(ctx, query,
		name,
		email,
		passwordHash,
		pinHash,
	).Scan(&userID)

	return userID, err
}

func (r *Repository) CreateAccountTx(
	ctx context.Context,
	tx pgx.Tx,
	userID int64,
	accountNumber string,
) error {
	query := `INSERT INTO accounts(user_id, account_number, balance) VALUES($1,$2,0)`

	_, err := tx.Exec(ctx, query, userID, accountNumber)

	return err
}
