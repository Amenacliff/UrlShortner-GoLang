package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"url_shortner/constants"
	"url_shortner/routes"
	"url_shortner/util"
)

func main() {

	app := fiber.New()

	app.Use(cors.New())

	envData, err := util.GetEnvData()

	if err != nil {
		return
	}

	mongoDBClient := util.ConnectToDb()

	routes.SetUpAllRoutes(app, mongoDBClient)

	log.Print(fmt.Sprintf("Connected to Port : %s", envData[constants.PORT]))
	log.Fatal(app.Listen(":" + envData[constants.PORT]))

}
