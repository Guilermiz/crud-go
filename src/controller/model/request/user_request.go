package request

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=4,containsany=!@#$%&*?"`
	Name      string `json:"name" binding:"required,min=3,max=50"`
	Age       int    `json:"age" binding:"required,numeric,min=18"`
}
