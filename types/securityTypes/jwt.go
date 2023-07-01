package securityTypes

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.MapClaims
}
