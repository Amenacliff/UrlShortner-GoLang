package shortUrl

type CreateShortUrlTO struct {
	LongURL    string `json:"longURL" bson:"longURL"`
	Passworded bool   `json:"passworded" bson:"passworded"`
	Password   string `json:"password" bson:"password"`
}
