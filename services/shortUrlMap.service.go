package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/genericMongo"
	"url_shortner/models"
)

type ShortUrlsService struct {
	Collection   *mongo.Collection
	GenericMongo *genericMongo.GenericMongo[models.ShortURLMap]
}

func (service *ShortUrlsService) Create(urlId string, longUrl string, passworded bool, password string, expiryTime int) error {
	newShortMapService := models.ShortURLMap{
		UrlId:        urlId,
		LongURL:      longUrl,
		Passworded:   passworded,
		Password:     password,
		ExpiryDate:   expiryTime,
		NumberOfHits: 0,
	}

	_, error := service.Collection.InsertOne(context.Background(), newShortMapService)

	if error != nil {
		return error
	}

	return nil

}
