package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"url_shortner/models"
)

type ShortUrlMapService struct {
	collection *mongo.Collection
}

func (service *ShortUrlMapService) Create(urlId string, longUrl string, passworded bool, password string, expiryTime int) error {
	newShortMapService := models.ShortURLMap{
		UrlId:        urlId,
		LongURL:      longUrl,
		Passworded:   passworded,
		Password:     password,
		ExpiryDate:   expiryTime,
		NumberOfHits: 0,
	}

	_, error := service.collection.InsertOne(context.Background(), newShortMapService)

	if error != nil {
		return error
	}

	return nil

}
