package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetMyAccount struct {
	ID            uuid.UUID `json:"id" example:"b2c1d9a2-7c2e-4b8b-a7a2-9e4a3f9d2e11"`
	UserID        uuid.UUID `json:"user_id" example:"b2c1d9a2-7c2e-4b8b-a7a2-9e4a3f9d2e11"`
	AccountNumber string    `json:"account_number" example:"BANK2025000001"`
	Balance       int64     `json:"balance" example:"5000"`
	Status        string    `json:"status" example:"active"`
	CreatedAt     time.Time `json:"created_at"`
}

type GetMyAccountUnathorised struct {
	Message string `json:"message" example:"invalid or expired token"`
}

type GetMyAccountServerError struct {
	Message string `json:"message" example:"Internal server error"`
}

type DepositRequest struct {
	AccountID uuid.UUID `json:"account_id" binding:"required" example:"b2c1d9a2-7c2e-4b8b-a7a2-9e4a3f9d2e11 (accountid value from getmyaccount)"`
	Amount    int64     `json:"amount" binding:"required,gt=0" example:"1000"`
}

type DepositResponse struct {
	Message string `json:"message" example:"deposit successful"`
}

type DepositResponseClientError struct {
	Message string `json:"message" example:"Bad Request"`
}

type DepositResponseServerError struct {
	Message string `json:"message" example:"Internal Server Error"`
}
