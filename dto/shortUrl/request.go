package shortUrl

type CreateShortUrlTO struct {
	UserId     string `json:"_id" bson:"_id"`
	LongURL    string `json:"longURL" bson:"longURL"`
	Passworded bool   `json:"passworded" bson:"passworded"`
	Password   string `json:"password" bson:"password"`
}
