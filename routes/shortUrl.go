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

	shortUrlController = controller
}

func (shortUrlRoute *ShortUrlRoutes) SetUpRoutes(app *fiber.App, shortUrlController *controller.ShortURLController, dbClient *mongo.Database) {

	shortUrlMapCollection := dbClient.Collection(constants.SHORT_URL_MAP_COLLECTION_NAME)

	shortUrlRoute.Init(shortUrlMapCollection)

	shortURLRoute := app.Group("/shortUrl")
	shortURLRoute.Get("/create", shortUrlController.Create)
}
