package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type MyCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

// key 类型为[]byte
var mySigningKey = []byte(viper.GetString("jwt.key"))

func GetToken(id uint, name string) (string, error) {
	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.accessTokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ikun",
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)

}

func ParseToken(tokenString string) (MyCustomClaims, error) {
	var claims MyCustomClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	return claims, err
}
