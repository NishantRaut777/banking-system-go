package dto

import "github.com/google/uuid"

type TransferRequest struct {
	FromAccountID uuid.UUID `json:"from_account_id" binding:"required" example:"b2c1d9a2-7c2e-4b8b-a7a2-9e4a3f9d2e11 (accountid value from getmyaccount)"`
	ToAccountID   uuid.UUID `json:"to_account_id" binding:"required" example:"b2c1d9a2-7c2e-4b8b-a7a2-9e4a3f9d2e11 (accountid value of receiver account)"`
	Amount        int64     `json:"amount" binding:"required,gt=0" example:"200"`
}

type TransferResponse struct {
	Message string `json:"message" example:"transfer successful"`
}

type TransferResponseClientError struct {
	Message string `json:"message" example:"Bad Request"`
}

type TransferResponseServerError struct {
	Message string `json:"message" example:"Internal Server Error"`
}
