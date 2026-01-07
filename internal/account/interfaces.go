package account

import (
	"context"

	"github.com/NishantRaut777/banking-api/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type AccountService interface {
	GetMyAccounts(ctx context.Context, userID uuid.UUID) ([]models.Account, error)

	GetAccountByID(ctx context.Context, userID, accountID uuid.UUID) (*models.Account, error)

	Deposit(ctx context.Context,
		userID uuid.UUID,
		accountID uuid.UUID,
		amount int64,
	) error
}

type AccountRepository interface {
	FindByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]models.Account, error)

	FindByIDAndUserID(
		ctx context.Context,
		accountID, userID uuid.UUID,
	) (*models.Account, error)

	GetAccountForUpdate(
		ctx context.Context,
		tx pgx.Tx,
		accountID uuid.UUID,
	) (uuid.UUID, int64, error)

	UpdateBalanceTx(
		ctx context.Context,
		tx pgx.Tx,
		accountID uuid.UUID,
		newBalance int64,
	) error

	InsertTransactionTx(
		ctx context.Context,
		tx pgx.Tx,
		accountID uuid.UUID,
		txType string,
		amount int64,
		status string,
		balanceBefore int64,
		balanceAfter int64,
	) error
}
