package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var signingKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id uint, name string) (string, error) {
	claims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

func ParseToken(token string) (JwtCustomClaims, error) {
	claims := JwtCustomClaims{}
	gToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	// fmt.Println(err)
	// fmt.Println(gToken.Valid)
	if err == nil && !gToken.Valid {
		// global.Logger.Error("Invalid Token")

		err = errors.New("Invalid Token")
	}
	return claims, err
}

func IsTokenValid(token string) bool {
	_, err := ParseToken(token)
	if err != nil {
		return false
	}
	return true
}
