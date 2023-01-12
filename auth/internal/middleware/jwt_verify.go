package middleware

import (
	"io/ioutil"
	"log"

	"github.com/golang-jwt/jwt"
)

func ValidateJwt(signedToken string) string {
	// Load the RSA public key from a file
	publicKey, err := ioutil.ReadFile("public.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the signed JWT and verify it with the RSA public key
	token, err := jwt.ParseWithClaims(signedToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return rsaPublicKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if token.Valid {
		return "TOKEN IS VALID"
	} else {
		return "TOKEN IS INVALID"
	}
}
