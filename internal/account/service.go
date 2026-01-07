package account

import (
	"context"
	"errors"

	"github.com/NishantRaut777/banking-api/internal/database"
	"github.com/NishantRaut777/banking-api/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo AccountRepository
}

func NewService(repo AccountRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetMyAccounts(
	ctx context.Context,
	userID uuid.UUID,
) ([]models.Account, error) {

	return s.repo.FindByUserID(ctx, userID)
}

func (s *Service) GetAccountByID(
	ctx context.Context,
	userID, accountID uuid.UUID,
) (*models.Account, error) {

	account, err := s.repo.FindByIDAndUserID(ctx, accountID, userID)
	if err != nil {
		return nil, errors.New("account not found or not authorized")
	}

	return account, nil
}

func (s *Service) Deposit(
	ctx context.Context,
	userID uuid.UUID,
	accountID uuid.UUID,
	amount int64,
) error {

	tx, err := database.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	accountUserID, balanceBefore, err := s.repo.GetAccountForUpdate(ctx, tx, accountID)
	if err != nil {
		return err
	}

	// allow deposit only if logged in user is trying to deposit money in own account
	if accountUserID != userID {
		return errors.New("unauthorized account access")
	}

	balanceAfter := balanceBefore + amount

	if err := s.repo.UpdateBalanceTx(ctx, tx, accountID, balanceAfter); err != nil {
		return err
	}

	if err := s.repo.InsertTransactionTx(
		ctx,
		tx,
		accountID,
		"deposit",
		amount,
		"success",
		balanceBefore,
		balanceAfter,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *Service) Withdraw(
	ctx context.Context,
	userID uuid.UUID,
	accountID uuid.UUID,
	amount int64,
) error {
	tx, err := database.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	accountUserID, balanceBefore, err := s.repo.GetAccountForUpdate(ctx, tx, accountID)
	if err != nil {
		return err
	}

	// allow withdraw only if logged in user is trying to deposit money in own account
	if accountUserID != userID {
		return errors.New("unauthorized account access")
	}

	if balanceBefore < amount {
		return errors.New("insufficient balance")
	}

	balanceAfter := balanceBefore - amount

	if err := s.repo.UpdateBalanceTx(ctx, tx, accountID, balanceAfter); err != nil {
		return err
	}

	if err := s.repo.InsertTransactionTx(
		ctx,
		tx,
		accountID,
		"withdraw",
		amount,
		"success",
		balanceBefore,
		balanceAfter,
	); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
