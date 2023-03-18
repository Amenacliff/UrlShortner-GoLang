package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/constants"
	"url_shortner/controller"
	"url_shortner/genericMongo"
	"url_shortner/services"
)

func SetUpAllRoutes(app *fiber.App, mongoDBClient *mongo.Client) {

	dataBase := mongoDBClient.Database(constants.DATABASE_NAME)
	shortUrlsCollection := dataBase.Collection(constants.SHORT_URL_MAP_COLLECTION_NAME)
	genericMongoose := genericMongo.GenericMongo{
		Collection: &mongo.Collection{},
	}
	shortUrlRoutes := ShortUrlRoutes{}
	shortUrlsService := &services.ShortUrlsService{
		Collection:   shortUrlsCollection,
		GenericMongo: &genericMongoose,
	}
	shortUrlController := &controller.ShortURLController{
		ShortUrlMapService: *shortUrlsService,
	}
	shortUrlRoutes.SetUpRoutes(app, shortUrlController, dataBase)

}
