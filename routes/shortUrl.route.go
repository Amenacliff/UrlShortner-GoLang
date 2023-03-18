package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/constants"
	"url_shortner/controller"
	"url_shortner/genericMongo"
	"url_shortner/services"
)

type ShortUrlRoutes struct {
	shortUrlMapService services.ShortUrlsService
	shortUrlController controller.ShortURLController
}

var shortUrlController controller.ShortURLController

func (shortUrlRoute *ShortUrlRoutes) Init(shortUrlMapCollection *mongo.Collection) {
	genericMongoClient := &genericMongo.GenericMongo{
		Collection: shortUrlMapCollection,
	}
	shortUrlMapService := services.ShortUrlsService{
		Collection:   shortUrlMapCollection,
		GenericMongo: genericMongoClient,
	}
	controller := controller.ShortURLController{
		ShortUrlMapService: shortUrlMapService,
	}

	shortUrlRoute.shortUrlController = controller
}

func (shortUrlRoute *ShortUrlRoutes) SetUpRoutes(app *fiber.App, dbClient *mongo.Database) {

	shortUrlMapCollection := dbClient.Collection(constants.SHORT_URL_MAP_COLLECTION_NAME)

	shortUrlRoute.Init(shortUrlMapCollection)

	shortURLRoute := app.Group("/shortUrl")
	shortURLRoute.Get("/create", shortUrlRoute.shortUrlController.Create)
}
