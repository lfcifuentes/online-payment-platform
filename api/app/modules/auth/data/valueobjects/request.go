package valueobjects

type LoginParams struct {
	Email    string `form:"email" validate:"required,email" documentation:"Email"`
	Password string `form:"password" validate:"required,password" documentation:"Contraseña"`
}

type RegisterParams struct {
	Name     string `form:"name" validate:"required" documentation:"Nombre"`
	Email    string `form:"email" validate:"required,email" documentation:"Email"`
	Password string `form:"password" validate:"required,password" documentation:"Contraseña"`
}
