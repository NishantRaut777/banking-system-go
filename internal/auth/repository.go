package auth

import (
	"context"

	"github.com/NishantRaut777/banking-system-go/internal/database"
	"github.com/NishantRaut777/banking-system-go/internal/models"
	"github.com/google/uuid"
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
) (uuid.UUID, error) {
	var userID uuid.UUID

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
	userID uuid.UUID,
	accountNumber string,
) error {
	query := `INSERT INTO accounts(user_id, account_number, balance) VALUES($1,$2,0)`

	_, err := tx.Exec(ctx, query, userID, accountNumber)

	return err
}

func (r *Repository) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	query := `SELECT id, email, password_hash,status FROM users WHERE email = $1`

	var user models.User

	err := database.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Status)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByID(
	ctx context.Context,
	userID uuid.UUID,
) (*models.User, error) {
	query := `SELECT id, name, email, status, created_at FROM users WHERE id = $1`

	var user models.User

	err := database.DB.QueryRow(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Status,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
