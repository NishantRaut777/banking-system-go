package dto

import "github.com/google/uuid"

type WithdrawRequest struct {
	AccountID uuid.UUID `json:"account_id" binding:"required" example:"b2c1d9a2-7c2e-4b8b-a7a2-9e4a3f9d2e11 (accountid value from getmyaccount)"`
	Amount    int64     `json:"amount" binding:"required,gt=0" example:"500"`
}

type WithdrawResponse struct {
	Message string `json:"message" example:"withdraw successful"`
}

type WithdrawResponseClientError struct {
	Message string `json:"message" example:"Bad Request"`
}

type WithdrawResponseServerError struct {
	Message string `json:"message" example:"Internal Server Error"`
}
