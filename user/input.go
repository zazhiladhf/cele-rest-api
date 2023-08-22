package user

type RegisterUserInput struct {
	Name        string `json:"name" binding:"required"`
	NoHandphone string `json:"no_handphone" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
