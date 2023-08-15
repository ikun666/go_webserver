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

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// key 类型为[]byte
var mySigningKey = []byte(viper.GetString("jwt.key"))

func GetToken(id uint, name string) (Tokens, error) {
	// Create claims with multiple fields populated
	accessClaims := MyCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.accessTokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ikun",
			Subject:   "accessToken",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(mySigningKey)
	if err != nil {
		return Tokens{}, err
	}
	refreshClaims := MyCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.refreshTokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "ikun",
			Subject:   "refreshToken",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		return Tokens{}, err
	}
	token := Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	return token, err

}

func ParseToken(accessToken, refreshToken string) (MyCustomClaims, bool, error) {
	var claims MyCustomClaims
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	// if err != nil {
	// 	fmt.Println("access token parse err")
	// 	return claims, false, err
	// }
	// if !atoken.Valid {
	// 	fmt.Println("access token invalid")
	// 	rtoken, err := jwt.ParseWithClaims(refreshToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 		return mySigningKey, nil
	// 	})
	// 	if err != nil {
	// 		fmt.Println("refresh token parse err")
	// 		return claims, false, err
	// 	}
	// 	if !rtoken.Valid {
	// 		fmt.Println("refresh token invalid")
	// 		return claims, false, err
	// 	}
	// }

	if err != nil {
		_, err := jwt.ParseWithClaims(refreshToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err == nil {
			return claims, true, nil
		}
	}

	return claims, false, err
}
