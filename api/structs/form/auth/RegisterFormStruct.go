package auth

type RegisterFormStruct struct {
	Name          string `form:"name" validate:"required"`
	Surname       string `form:"surname" validate:"required"`
	Email         string `form:"email" validate:"required,email"`
	Password      string `form:"password" validate:"required,min=4"`
	PasswordRetry string `form:"password_retry" validate:"required,eqfield=Password,min=4"`
}
