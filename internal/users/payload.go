package users

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type RegisterResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
