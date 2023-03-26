package userDTO

type CreateUserDTO struct {
	EmailAddress string `json:"email" bson:"email"`
	Password     string `json:"password"`
}
