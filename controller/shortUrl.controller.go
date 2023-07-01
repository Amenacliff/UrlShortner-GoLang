package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"url_shortner/dto/shortUrl"
	"url_shortner/models"
	"url_shortner/services"
	"url_shortner/util"
)

type ShortURLController struct {
	ShortUrlMapService services.ShortUrlsService
	UserService        services.UserService
}

func (controller *ShortURLController) Create(ctx *fiber.Ctx) error {
	createShortUrl := shortUrl.CreateShortUrlTO{}
	err := ctx.BodyParser(&createShortUrl)

	if err != nil {
		log.Println(err.Error())
		return util.GenerateResponse(ctx, "Something Went Wrong", false, "Something Went Wrong")
	}

	userObjectId, errGetUserObjectId := primitive.ObjectIDFromHex(createShortUrl.UserId)

	if errGetUserObjectId != nil {
		log.Println(errGetUserObjectId.Error())
		return util.GenerateResponse(ctx, "Something Went Wrong", false, "Something Went Wrong")
	}

	_, errGetUser := controller.UserService.GenericMongo.FindOne(util.GetFieldBsonTag[models.User]([]models.User{{ID: userObjectId}}), []any{userObjectId})

	if errGetUser != nil {
		return util.GenerateResponse(ctx, "User not Found", false, "Something Went Wrong")
	}

	return util.GenerateResponse(ctx, "", true, "")
}
