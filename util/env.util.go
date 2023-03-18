package util

import (
	"github.com/joho/godotenv"
	"log"
)

func GetEnvData() (map[string]string, error) {
	var myEnv map[string]string
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		return myEnv, err
	} else {
		newlyEnv, errFromRead := godotenv.Read()
		myEnv = newlyEnv
		if errFromRead != nil {
			log.Println(errFromRead.Error())
			return myEnv, errFromRead
		}
	}

	return myEnv, nil
}
