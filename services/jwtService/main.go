package jwtService

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strconv"
	"time"
	"url_shortner/constants"
	"url_shortner/types/securityTypes"
	"url_shortner/util"
)

func CreateToken(userId string) (string, error) {
	signingKey, errFromEnv := util.GetEnv(constants.JWTSigningKey)
	if errFromEnv != nil {
		log.Println(errFromEnv)
		return "", errFromEnv
	}

	expirationTime := time.Now().Add(time.Hour * 168).Unix()
	jwtClaims := securityTypes.JWTClaims{
		UserId: userId,
		MapClaims: jwt.MapClaims{
			strconv.FormatInt(expirationTime, 10): expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {
	signingKey, errFromEnv := util.GetEnv(constants.JWTSigningKey)
	if errFromEnv != nil {
		log.Println(errFromEnv)
		return false, errFromEnv
	}

	token, errorFromParseJWT := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if errorFromParseJWT != nil {
		log.Println(errorFromParseJWT)
		return false, errorFromParseJWT
	}

	return token.Valid, nil
}

func GetUserId(tokenString string) (string, error) {
	signingKey, errFromEnv := util.GetEnv(constants.JWTSigningKey)
	if errFromEnv != nil {
		log.Println(errFromEnv)
		return "", errFromEnv
	}

	userClaims := &securityTypes.JWTClaims{}

	token, errorFromParseClaims := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if errorFromParseClaims != nil {
		log.Println(errorFromParseClaims.Error())
		return "", errorFromParseClaims
	}

	if token.Valid == false {
		return "", errors.New("invalid token ")
	}

	return userClaims.UserId, nil
}
