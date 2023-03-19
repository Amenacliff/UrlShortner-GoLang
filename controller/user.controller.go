package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	userDTO "url_shortner/dto/user"
	"url_shortner/models"
	"url_shortner/services"
	"url_shortner/util"
)

type UserController struct {
	UserService *services.UserService
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	createUser := userDTO.CreateUserDTO{}
	err := ctx.BodyParser(&createUser)

	if err != nil {
		log.Println(err.Error())
		return util.GenerateResponse[interface{}](ctx, nil, false, "Request Body Not Found")
	}

	userSearch := models.User{Email: createUser.EmailAddress}

	user, errGetUser := controller.UserService.GenericMongo.FindOne(util.StructToMap(userSearch))

	if errGetUser != nil {
		return util.GenerateResponse[interface{}](ctx, nil, false, errGetUser.Error())
	}

	log.Println(user)

	return nil

}
