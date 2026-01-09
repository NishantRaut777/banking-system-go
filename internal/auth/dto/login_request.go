package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@gmail.com"`
	Password string `json:"password" binding:"required" example:"StrongPass123"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type LoginResponseClientError struct {
	Message string `json:"message" example:"Bad Request"`
}

type LoginResponseServerError struct {
	Message string `json:"message" example:"Internal Server Error"`
}
