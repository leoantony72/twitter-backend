package services

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt"
)

func (u *userUseCase) GetTokenClaims(token string) (*jwt.Token, jwt.MapClaims, error) {
	// Load the RSA public key from a file
	publicKey, err := ioutil.ReadFile("../public.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	refresh_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexepcted signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	claims, ok := refresh_token.Claims.(jwt.MapClaims)
	if ok && refresh_token.Valid {
		return refresh_token, claims, nil
	}
	return nil, nil, err
}
