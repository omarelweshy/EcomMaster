package form

type RegistrationForm struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required,min=8"`
	Email     string `json:"email" binding:"required,email"`
}
