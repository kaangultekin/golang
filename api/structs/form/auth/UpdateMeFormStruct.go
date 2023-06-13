package auth

type UpdateMeFormStruct struct {
	Name    string `form:"name" validate:"required"`
	Surname string `form:"surname" validate:"required"`
	Email   string `form:"email" validate:"required,email"`
}
