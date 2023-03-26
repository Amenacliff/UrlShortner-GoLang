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

func GetFieldBsonTag[T any](object T) []string {

	typeReflect := reflect.TypeOf(object)
	structValue := reflect.ValueOf(object)

	allBsonFields := make([]string, 0)

	for i := 0; i < typeReflect.NumField(); i++ {

		if structValue.Field(i).IsZero() == false {
			allBsonFields = append(allBsonFields, typeReflect.Field(i).Tag.Get("bson"))
		}
	}

	return allBsonFields

}
