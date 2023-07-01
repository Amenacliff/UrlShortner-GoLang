package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"log"
	"url_shortner/constants"
	"url_shortner/routes"
	"url_shortner/util"
)

func main() {
	envData, err := util.GetEnvData()

	encryptKey, err := util.GetEnv(constants.ENCRYPT_KEY)
	app := fiber.New()

	app.Use(cors.New())

	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptKey,
	}))

	if err != nil {
		return
	}

	mongoDBClient := util.ConnectToDb()

	routes.SetUpAllRoutes(app, mongoDBClient)

	log.Print(fmt.Sprintf("Connected to Port : %s", envData[constants.PORT]))
	log.Fatal(app.Listen(":" + envData[constants.PORT]))

}
