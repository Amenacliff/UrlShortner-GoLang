package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/genericMongo"
	"url_shortner/models"
)

type UserService struct {
	Collection   *mongo.Collection
	GenericMongo *genericMongo.GenericMongo
}

func (userSer *UserService) Create(emailAddress, passwordHash string) error {
	newUser := models.User{
		Email:        "",
		PasswordHash: "",
	}

	_, err := userSer.Collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		return err
	}

	return nil
}
