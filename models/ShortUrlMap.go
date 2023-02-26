package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortURLMap struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	UrlId        string             `json:"urlId" bson:"urlId"`
	LongURL      string             `json:"longURL" bson:"longURL"`
	Passworded   bool               `json:"passworded" bson:"passworded"`
	Password     string             `json:"password" bson:"password"`
	ExpiryDate   int                `json:"expiryDate" bson:"expiryDate"`
	NumberOfHits int                `json:"numberOfHits" bson:"numberOfHits"`
}
