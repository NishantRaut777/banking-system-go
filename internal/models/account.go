package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	AccountNumber string    `json:"account_number"`
	Balance       int64     `json:"balance"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

type DepositRequest struct {
	AccountID uuid.UUID `json:"account_id" binding:"required"`
	Amount    int64     `json:"amount" binding:"required,gt=0"`
}

type WithdrawRequest struct {
	AccountID uuid.UUID `json:"account_id" binding:"required"`
	Amount    int64     `json:"amount" binding:"required,gt=0"`
}
