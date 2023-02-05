package request

type ValidateRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type ValidateLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
