package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func StructToMap(obj interface{}) (newMap map[any]any) {
	data, err := bson.Marshal(obj)

	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &newMap)
	return
}

func GetFieldBsonTag[T any](objects []T) []string {

	allBsonFields := make([]string, 0)

	for _, object := range objects {
		typeReflect := reflect.TypeOf(object)
		structValue := reflect.ValueOf(object)

		for i := 0; i < structValue.NumField(); i++ {

			if structValue.Field(i).IsZero() == false {
				allBsonFields = append(allBsonFields, typeReflect.Field(i).Tag.Get("bson"))
			}
		}
	}

	return allBsonFields

}
