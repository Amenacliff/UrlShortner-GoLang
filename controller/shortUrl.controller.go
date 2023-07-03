package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
	"url_shortner/constants"
	"url_shortner/dto/shortUrl"
	"url_shortner/models"
	"url_shortner/services"
	"url_shortner/services/jwtService"
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

	cookie := ctx.Cookies(constants.CookieKey)

	userId, errGetId := jwtService.GetUserId(cookie)

	if errGetId != nil {
		log.Println(errGetId.Error())
		return util.GenerateResponse(ctx, "Unauthorized", false, "Not Logged in ")

	}

	userObjectId, errGetUserObjectId := primitive.ObjectIDFromHex(userId)

	if errGetUserObjectId != nil {
		log.Println(errGetUserObjectId.Error())
		return util.GenerateResponse(ctx, "Something Went Wrong", false, "Something Went Wrong")
	}

	_, errGetUser := controller.UserService.GenericMongo.FindOne(util.GetFieldBsonTag[models.User]([]models.User{{ID: userObjectId}}), []any{userObjectId})

	if errGetUser != nil {
		return util.GenerateResponse(ctx, "User not Found", false, "Something Went Wrong")
	}

	urlId := uuid.New().String()[0:5]

	errCreate := controller.ShortUrlMapService.Create(urlId, createShortUrl.LongURL, createShortUrl.Passworded, createShortUrl.Password, int(time.Now().Add(time.Hour*24*30).UnixMilli()), userId)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return util.GenerateResponse(ctx, "Something Went Wrong", false, errCreate.Error())
	}

	return util.GenerateResponse(ctx, urlId, true, "Generated Short URL successfully")

}
