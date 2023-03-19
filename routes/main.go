package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/constants"
)

func SetUpAllRoutes(app *fiber.App, mongoDBClient *mongo.Client) {

	dataBase := mongoDBClient.Database(constants.DATABASE_NAME)

	shortUrlRoutes := ShortUrlRoutes{}
	userRoutes := UserRoute{}

	shortUrlRoutes.SetUpRoutes(app, dataBase)
	userRoutes.SetUpRoute(app, dataBase)

}
