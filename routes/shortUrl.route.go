package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/constants"
	"url_shortner/controller"
	"url_shortner/genericMongo"
	"url_shortner/models"
	"url_shortner/services"
)

type ShortUrlRoutes struct {
	shortUrlMapService services.ShortUrlsService
	shortUrlController controller.ShortURLController
}

var shortUrlController controller.ShortURLController

func (shortUrlRoute *ShortUrlRoutes) Init(shortUrlMapCollection *mongo.Collection, userCollection *mongo.Collection) {
	genericShortUrlMongoClient := &genericMongo.GenericMongo[models.ShortURLMap]{
		Collection: shortUrlMapCollection,
	}
	shortUrlMapService := services.ShortUrlsService{
		Collection:   shortUrlMapCollection,
		GenericMongo: genericShortUrlMongoClient,
	}

	genericUserCollection := &genericMongo.GenericMongo[models.User]{
		Collection: userCollection,
	}

	userService := services.UserService{
		Collection:   userCollection,
		GenericMongo: genericUserCollection,
	}

	controller := controller.ShortURLController{
		ShortUrlMapService: shortUrlMapService,
		UserService:        userService,
	}

	shortUrlRoute.shortUrlController = controller
}

func (shortUrlRoute *ShortUrlRoutes) SetUpRoutes(app *fiber.App, dbClient *mongo.Database) {

	shortUrlMapCollection := dbClient.Collection(constants.SHORT_URL_MAP_COLLECTION_NAME)

	userCollection := dbClient.Collection(constants.USER_COLLECTION_NAME)

	shortUrlRoute.Init(shortUrlMapCollection, userCollection)

	shortURLRoute := app.Group("/shortUrl")
	shortURLRoute.Post("/create", shortUrlRoute.shortUrlController.Create)
}
