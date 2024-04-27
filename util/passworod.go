package util

import (
	"github.com/joomcode/errorx"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errorx.InternalError.Wrap(err, "failed to hash password:")
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func HashToken(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", errorx.InternalError.Wrap(err, "failed to hash token")
	}
	return string(hashedToken), nil
}

func CheckToken(token, hashedToken string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(token))	
}
