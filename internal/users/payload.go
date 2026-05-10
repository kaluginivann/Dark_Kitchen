package users

type RegisterInput struct {
	Body RegisterRequest
}

type RegisterRequest struct {
	Username string `json:"username" minLength:"3"`
	Password string `json:"password" minLength:"8"`
	Email    string `json:"email" format:"email"`
}

type RegisterOutput struct {
	Body RegisterResponse
}

type RegisterResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
