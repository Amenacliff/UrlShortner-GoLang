package genericMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
import appContext "context"

type GenericMongo struct {
	Collection *mongo.Collection
}

func (genericMongoose *GenericMongo) FindOne(keysAndValues map[any]any) (any, error) {

	var document any

	err := genericMongoose.Collection.FindOne(appContext.Background(), keysAndValues).Decode(&document)

	if err != nil {
		return document, err
	}

	return document, nil

}

func (genericMongoose *GenericMongo) Find(keysAndValues map[any]any) ([]any, error) {
	var allDocuments []any

	results, err := genericMongoose.Collection.Find(appContext.Background(), keysAndValues)

	if err != nil {
		return allDocuments, err
	}

	errParseResults := results.All(appContext.Background(), &allDocuments)

	if errParseResults != nil {
		return allDocuments, errParseResults
	}

	return allDocuments, nil
}

func (genericMongoose *GenericMongo) GetAll() ([]any, error) {

	var allDocuments []any

	results, err := genericMongoose.Collection.Find(appContext.Background(), bson.D{})

	if err != nil {
		return allDocuments, err
	}

	errParseResults := results.All(appContext.Background(), &allDocuments)

	if errParseResults != nil {
		return allDocuments, errParseResults
	}

	return allDocuments, nil

}
