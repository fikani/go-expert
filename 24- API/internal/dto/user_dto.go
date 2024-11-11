package dto

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtTokenInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtTokenOutput struct {
	Token string `json:"token"`
}
