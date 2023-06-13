package auth

type UpdatePasswordFormStruct struct {
	OldPassword   string `form:"old_password" validate:"required,min=4"`
	Password      string `form:"password" validate:"required,min=4"`
	PasswordRetry string `form:"password_retry" validate:"required,eqfield=Password,min=4"`
}
