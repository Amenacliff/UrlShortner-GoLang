package userDTO

type CreateUserDTO struct {
	EmailAddress string `json:"emailAddress"`
	Password     string `json:"password"`
}
