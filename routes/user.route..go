package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/controller"
	"url_shortner/genericMongo"
	"url_shortner/services"
)

type UserRoute struct {
	UserService    *services.UserService
	UserController *controller.UserController
}

func (route *UserRoute) Init(userCollection *mongo.Collection) {
	genericMongoClient := &genericMongo.GenericMongo{
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

func (route *UserRoute) SetUpRoute(app *fiber.App, userCollection *mongo.Collection) {
	userGroup := app.Group("/user")
	route.Init(userCollection)
	userGroup.Post("/create", route.UserController.Create)
}
