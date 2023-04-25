package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateAccessToken(name, Id string) (string, error) {
	privateKey, err := ioutil.ReadFile("private.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaprivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"name": name,
		"ID":   Id,
		"exp":  time.Now().Add(10 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(rsaprivateKey)
	if err != nil {
		return "", errors.New("failed to create Token")
	}
	return tokenString, nil
}

func GenerateRefreshToken(name string, Id string) (string, error) {
	privateKey, err := ioutil.ReadFile("private.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaprivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"name": name,
		"ID":   Id,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	refreshTokenString, err := token.SignedString(rsaprivateKey)
	if err != nil {
		return "", errors.New("failed to create Token")
	}
	return refreshTokenString, nil
}
