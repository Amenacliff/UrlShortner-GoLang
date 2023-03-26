package genericMongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)
import appContext "context"

type GenericMongo[T any] struct {
	Collection *mongo.Collection
}

func (genericMongoose *GenericMongo[T]) FindOne(keys []string, values []any) (T, error) {

	log.Println(keys, values)

	var document T

	var bsonKeyAndValue = bson.M{}

	for i, key := range keys {
		bsonKeyAndValue[key] = values[i]
	}

	err := genericMongoose.Collection.FindOne(appContext.Background(), bsonKeyAndValue).Decode(&document)

	if err != nil {
		return document, err
	}

	return document, nil

}

func (genericMongoose *GenericMongo[T]) Find(keysAndValues T) ([]T, error) {
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

func (genericMongoose *GenericMongo[T]) FindById(id primitive.ObjectID) (T, error) {
	filterObject := bson.D{{
		Key:   "_id",
		Value: id,
	}}

	var document T

	err := genericMongoose.Collection.FindOne(appContext.TODO(), filterObject).Decode(&document)

	if err != nil {
		return document, err
	}

	return document, nil

}
