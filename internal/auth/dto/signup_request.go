package dto

type SignupRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"johndoe@gmail.com"`
	Password string `json:"password" binding:"required,min=8" example:"StrongPass123"`
	Pin      string `json:"pin" binding:"required,len=4" example:"1234"`
}

type SignupResponse struct {
	Message string `json:"message" example:"user registered successfully"`
}

type SignupResponseClientError struct {
	Message string `json:"message" example:"Bad Request"`
}

type SignupResponseServerError struct {
	Message string `json:"message" example:"Internal Server Error"`
}
