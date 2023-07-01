package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
	"url_shortner/constants"
	userDTO "url_shortner/dto/user"
	"url_shortner/models"
	"url_shortner/services"
	"url_shortner/services/jwtService"
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

	user, errGetUser := controller.UserService.GenericMongo.FindOne(util.GetFieldBsonTag[models.User]([]models.User{{Email: createUser.EmailAddress}}), []any{createUser.EmailAddress})

	log.Println(user, createUser)

	if errGetUser != nil && errGetUser != mongo.ErrNoDocuments {
		return util.GenerateResponse[interface{}](ctx, nil, false, errGetUser.Error())
	}

	if strings.TrimSpace(strings.ToLower(user.Email)) == strings.TrimSpace(strings.ToLower(createUser.EmailAddress)) {
		return util.GenerateResponse[interface{}](ctx, nil, false, "User Already Exists")
	}

	hashedPassword, errHashPass := bcrypt.GenerateFromPassword([]byte(createUser.Password), 14)

	if errHashPass != nil {
		log.Println(errHashPass.Error())
		return util.GenerateResponse[interface{}](ctx, nil, false, errHashPass.Error())
	}

	userId, errCreateUser := controller.UserService.Create(createUser.EmailAddress, string(hashedPassword))

	if errCreateUser != nil {
		log.Println(errCreateUser.Error())
		return util.GenerateResponse[interface{}](ctx, nil, false, errCreateUser.Error())
	}

	jwtToken, errCreateToken := jwtService.CreateToken(userId)

	if errCreateToken != nil {
		log.Println(errCreateToken.Error())
	}

	if errCreateUser == nil {
		cookieDetails := fiber.Cookie{
			Name:     constants.CookieKey,
			Value:    jwtToken,
			Expires:  time.Now().Add(time.Hour * 168),
			HTTPOnly: false,
			Secure:   false,
		}

		ctx.Cookie(&cookieDetails)
	}

	return util.GenerateResponse[interface{}](ctx, "User Created Successfully", true, "")

}
