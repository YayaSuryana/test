package user

// struct yang di pakai untuk mapping inputan dari user
type RegisterUserInput struct{
	Name		string	`json:"name" binding:"required"`
	Occupation 	string	`json:"occupation" binding:"required"`
	Email		string	`json:"email" binding:"required,email"`
	Password	string	`json:"password" binding:"required"`
}

type LoginInput struct{
	Email		string	`json:"email" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

type CheckEmailAvailable struct {
	Email 		string `json:"email" binding:"required,email"`
}