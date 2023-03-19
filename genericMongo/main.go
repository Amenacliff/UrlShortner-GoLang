package genericMongo

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
import appContext "context"

type GenericMongo[T any] struct {
	Collection *mongo.Collection
}

func (genericMongoose *GenericMongo[T]) FindOne(keysAndValues map[any]any) (T, error) {

	var document T

	err := genericMongoose.Collection.FindOne(appContext.Background(), keysAndValues).Decode(&document)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return document, errors.New("Document not found")
		}
		return document, err
	}

	return document, nil

}

func (genericMongoose *GenericMongo[T]) Find(keysAndValues map[any]any) ([]T, error) {
	var allDocuments []T

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

func (genericMongoose *GenericMongo[T]) GetAll() ([]T, error) {

	var allDocuments []T

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
