package account

import (
	"context"
	"errors"
	"net/http"

	"github.com/NishantRaut777/banking-system-go/internal/database"
	"github.com/NishantRaut777/banking-system-go/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FindMyAccount(
	ctx context.Context,
	userID uuid.UUID,
) (*models.Account, error, int) {

	var status int
	status = http.StatusOK

	query := `SELECT id, user_id, account_number, balance, status, created_at FROM accounts WHERE user_id = $1`

	var account models.Account
	err := database.DB.QueryRow(ctx, query, userID).
		Scan(
			&account.ID,
			&account.UserID,
			&account.AccountNumber,
			&account.Balance,
			&account.Status,
			&account.CreatedAt,
		)

	if err != nil {
		status = http.StatusBadRequest
		return nil, err, status
	}

	return &account, nil, status
}

func (r *Repository) FindByIDAndUserID(
	ctx context.Context,
	accountID, userID uuid.UUID,
) (*models.Account, error) {

	query := `
		SELECT id, user_id, account_number, balance, status, created_at
		FROM accounts
		WHERE id = $1 AND user_id = $2
	`

	var a models.Account
	err := database.DB.QueryRow(ctx, query, accountID, userID).
		Scan(
			&a.ID,
			&a.UserID,
			&a.AccountNumber,
			&a.Balance,
			&a.Status,
			&a.CreatedAt,
		)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *Repository) GetAccountForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	accountID uuid.UUID,
) (uuid.UUID, int64, error) {

	var userID uuid.UUID
	var balance int64

	query := `SELECT user_id, balance FROM accounts WHERE id = $1 FOR UPDATE`

	err := tx.QueryRow(ctx, query, accountID).Scan(&userID, &balance)

	return userID, balance, err
}

func (r *Repository) UpdateBalanceTx(
	ctx context.Context,
	tx pgx.Tx,
	accountID uuid.UUID,
	newBalance int64,
) error {

	query := `UPDATE accounts SET balance = $1 WHERE id = $2`

	_, err := tx.Exec(ctx, query, newBalance, accountID)

	return err
}

func (r *Repository) InsertTransactionTx(
	ctx context.Context,
	tx pgx.Tx,
	accountID uuid.UUID,
	txType string,
	amount int64,
	status string,
	balanceBefore int64,
	balanceAfter int64,
) error {

	query := `
		INSERT INTO transactions
		(account_id, type, amount, status, balance_before, balance_after)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := tx.Exec(
		ctx,
		query,
		accountID,
		txType,
		amount,
		status,
		balanceBefore,
		balanceAfter,
	)

	return err
}

func (r *Repository) GetAccountsForUpdate(
	ctx context.Context,
	tx pgx.Tx,
	acc1 uuid.UUID,
	acc2 uuid.UUID,
) (map[uuid.UUID]struct {
	UserID  uuid.UUID
	Balance int64
}, error) {

	query := `SELECT id, user_id, balance FROM accounts WHERE id IN ($1,$2) ORDER BY id FOR UPDATE`

	rows, err := tx.Query(ctx, query, acc1, acc2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[uuid.UUID]struct {
		UserID  uuid.UUID
		Balance int64
	})

	for rows.Next() {
		var id, userID uuid.UUID
		var balance int64
		if err := rows.Scan(&id, &userID, &balance); err != nil {
			return nil, err
		}

		result[id] = struct {
			UserID  uuid.UUID
			Balance int64
		}{userID, balance}
	}

	if len(result) != 2 {
		return nil, errors.New("one or both accounts not found")
	}

	return result, nil
}
