package genericMongo

import "go.mongodb.org/mongo-driver/mongo"
import appContext "context"

type GenericMongo struct {
	collection mongo.Collection
}

func (genericMongoose *GenericMongo) FindOne(keysAndValues map[any]any) (any, error) {

	var document any

	err := genericMongoose.collection.FindOne(appContext.Background(), keysAndValues).Decode(&document)

	if err != nil {
		return document, err
	}

	return document, nil

}

func (genericMongoose *GenericMongo) FindAll(keysAndValues map[any]any) ([]any, error) {
	var allDocuments []any

	results, err := genericMongoose.collection.Find(appContext.Background(), keysAndValues)

	if err != nil {
		return allDocuments, err
	}

	errParseResults := results.All(appContext.Background(), &allDocuments)

	if errParseResults != nil {
		return allDocuments, errParseResults
	}

	return allDocuments, nil
}
