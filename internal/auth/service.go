package auth

import (
	"context"
	"errors"

	"github.com/NishantRaut777/banking-api/internal/database"
	"github.com/NishantRaut777/banking-api/internal/utils"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo AuthRepository
}

func NewService(repo AuthRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Signup(
	ctx context.Context,
	name, email, password, pin string,
) error {
	passwordHash, err := utils.HashString(password)
	if err != nil {
		return err
	}

	pinHash, err := utils.HashString(pin)
	if err != nil {
		return err
	}

	tx, err := database.DB.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	// create user in "users"
	userID, err := s.repo.CreateUserTx(
		ctx,
		tx,
		name,
		email,
		passwordHash,
		pinHash,
	)

	if err != nil {
		return err
	}

	accountNumber := utils.GenerateAccountNumber(userID)

	// create account in "accounts"
	err = s.repo.CreateAccountTx(
		ctx,
		tx,
		userID,
		accountNumber,
	)

	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
