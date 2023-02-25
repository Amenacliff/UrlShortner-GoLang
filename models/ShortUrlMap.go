package models

type ShortURLMap struct {
	UrlId        string `json:"urlId" bson:"urlId"`
	LongURL      string `json:"longURL" bson:"longURL"`
	Passworded   string `json:"passworded" bson:"passworded"`
	Password     string `json:"password" bson:"password"`
	ExpiryDate   int    `json:"expiryDate" bson:"expiryDate"`
	NumberOfHits int    `json:"numberOfHits" bson:"numberOfHits"`
}
