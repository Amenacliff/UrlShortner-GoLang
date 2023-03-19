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

type UserRoute struct {
	UserService    *services.UserService
	UserController *controller.UserController
}

func (route *UserRoute) Init(userCollection *mongo.Collection) {
	genericMongoClient := &genericMongo.GenericMongo[models.User]{
		Collection: userCollection,
	}

	userService := &services.UserService{
		Collection:   userCollection,
		GenericMongo: genericMongoClient,
	}

	userController := &controller.UserController{
		UserService: userService,
	}

	route.UserService = userService
	route.UserController = userController
}

func (route *UserRoute) SetUpRoute(app *fiber.App, dataBase *mongo.Database) {
	userCollection := dataBase.Collection(constants.USER_COLLECTION_NAME)
	userGroup := app.Group("/user")
	route.Init(userCollection)
	userGroup.Post("/create", route.UserController.Create)
}
