package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type AuthService interface {
	Signup(
		ctx context.Context,
		name, email, password, pin string,
	) error
}

type AuthRepository interface {
	CreateUserTx(
		ctx context.Context,
		tx pgx.Tx,
		name, email, passwordHash, pinHash string,
	) (int64, error)

	CreateAccountTx(
		ctx context.Context,
		tx pgx.Tx,
		userID int64,
		accountNumber string,
	) error
}
