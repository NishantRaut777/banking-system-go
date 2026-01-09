package auth

import (
	"context"
	"errors"

	"github.com/NishantRaut777/banking-system-go/internal/database"
	"github.com/NishantRaut777/banking-system-go/internal/models"
	"github.com/NishantRaut777/banking-system-go/internal/utils"
	"github.com/google/uuid"
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

func (s *Service) Login(
	ctx context.Context,
	email, password string,
) (string, error) {

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	if user.Status != "active" {
		return "", errors.New("user account is not active")
	}

	if !utils.CompareHash(user.PasswordHash, password) {
		return "", errors.New("Invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetProfile(
	ctx context.Context,
	userID uuid.UUID,
) (*models.User, error) {

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
