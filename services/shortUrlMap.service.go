package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"log"
	"url_shortner/genericMongo"
	"url_shortner/models"
)

type ShortUrlsService struct {
	Collection   *mongo.Collection
	GenericMongo *genericMongo.GenericMongo[models.ShortURLMap]
}

func (service *ShortUrlsService) Create(urlId string, longUrl string, passworded bool, password string, expiryTime int, userId string) error {

	encryptedPassword := password

	if passworded == true {
		hashedPassword, errorGenPass := bcrypt.GenerateFromPassword([]byte(password), 14)

		if errorGenPass != nil {
			log.Println(errorGenPass.Error())
		}
		encryptedPassword = string(hashedPassword)
	}

	newShortMapService := models.ShortURLMap{
		ID:           primitive.NewObjectID(),
		UrlId:        urlId,
		LongURL:      longUrl,
		Passworded:   passworded,
		Password:     encryptedPassword,
		ExpiryDate:   expiryTime,
		NumberOfHits: 0,
		UserId:       userId,
	}

	_, error := service.Collection.InsertOne(context.Background(), newShortMapService)

	if error != nil {
		return error
	}

	return nil

}
