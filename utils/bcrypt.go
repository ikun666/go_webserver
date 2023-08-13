package utils

import "golang.org/x/crypto/bcrypt"

func Crypt(password string) (string, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(pwd), err
}

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
