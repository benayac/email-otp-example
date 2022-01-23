package model

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type RegisterResponse struct {
	Nonce string `json:"nonce"`
}

type OTPRequest struct {
	OTP   string `json:"otp"`
	Email string `json:"email"`
}

type OTPResponse struct {
	Status bool `json:"status"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type User struct {
	ID         uint
	Name       string
	Password   string
	Email      string
	Address    string
	IsVerified bool
}
