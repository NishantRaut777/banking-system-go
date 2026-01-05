package auth

import (
	"context"

	"github.com/NishantRaut777/banking-api/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type AuthService interface {
	Signup(
		ctx context.Context,
		name, email, password, pin string,
	) error

	Login(
		ctx context.Context,
		email, password string,
	) (string, error)
}

type AuthRepository interface {
	CreateUserTx(
		ctx context.Context,
		tx pgx.Tx,
		name, email, passwordHash, pinHash string,
	) (uuid.UUID, error)

	CreateAccountTx(
		ctx context.Context,
		tx pgx.Tx,
		userID uuid.UUID,
		accountNumber string,
	) error

	GetUserByEmail(
		ctx context.Context,
		email string,
	) (*models.User, error)
}
