package auth

type LoginFormStruct struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=4"`
}
