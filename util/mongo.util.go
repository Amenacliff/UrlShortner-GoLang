package util

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"url_shortner/constants"
)

func ConnectToDb() *mongo.Client {
	envData, errGetEnv := GetEnvData()
	if errGetEnv != nil {
		log.Println(errGetEnv.Error())
	}

	dataBaseURI := envData[constants.MONGO_URI]

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dataBaseURI))

	if err != nil {
		return nil
	}

	log.Println("DataBase Connected Successfully")

	return client

}
