package middleware

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"github.com/leoantony72/twitter-backend/auth/internal/utils"
)

func (u *UserMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		signedToken, err := ctx.Cookie("access-Token")
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Please Login"})
			ctx.Abort()
			return
		}
		refreshToken, err := ctx.Cookie("refresh-Token")
		if err != nil {
			ctx.JSON(400, gin.H{"message": "Please Login"})
			ctx.Abort()
			return
		}

		_, claims, err := ValidateJwt(signedToken)
		if err != nil {
			if err.Error() == "Token is expired" {
				check, user := u.CheckRefreshToken(refreshToken)
				if !check {
					ctx.AbortWithStatusJSON(401, gin.H{"msg": "Please Login"})
					return
				}
				accessToken, _ := utils.GenerateAccessToken(user.Username, user.Id)
				ctx.SetCookie("access-Token", accessToken, 3600, "/", "", false, true)
				ctx.Set("username", user.Username)
				ctx.Set("id", user.Id)
				ctx.Next()
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

func (u *UserMiddleware) CheckRefreshToken(refreshToken string) (bool, model.User) {
	user, _ := u.userUseCase.GetToken(refreshToken)
	return user.Token == refreshToken, user
}
