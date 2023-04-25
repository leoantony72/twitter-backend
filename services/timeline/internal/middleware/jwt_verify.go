package middleware

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (u *TimelineMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		signedToken, err := ctx.Cookie("access-Token")
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Please Login"})
			ctx.Abort()
			return
		}

		_, claims, err := ValidateJwt(signedToken)
		if err != nil {
			if err.Error() == "Token is expired" {
				ctx.AbortWithStatusJSON(401, gin.H{"msg": "Token expired"})
				return
			}
			ctx.JSON(500, gin.H{"message": "Something went wrong, Try Again"})
			ctx.Abort()
			return
		}
		ctx.Set("username", claims["name"].(string))
		ctx.Set("id", claims["ID"].(string))
		ctx.Next()

	}
}

func ValidateJwt(signedToken string) (*jwt.Token, jwt.MapClaims, error) {
	// Load the RSA public key from a file
	publicKey, err := ioutil.ReadFile("../public.pem")
	if err != nil {
		log.Fatal(err)
	}
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the signed JWT and verify it with the RSA public key
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexepcted signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return token, claims, nil
	}
	return nil, nil, err
}
