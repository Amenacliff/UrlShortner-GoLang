package util

import "encoding/json"

func StructToMap(obj interface{}) (newMap map[any]any) {
	data, err := json.Marshal(obj)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}
