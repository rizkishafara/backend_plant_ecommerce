package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	Email   string `json:"email"`
	Idreq   string `json:"idreq"`
	Expired string `json:"expiredtime"`
	jwt.StandardClaims
}

func Generatejwt(email, idreq, expiredtime string) (string, int) {
	claims := &jwtCustomClaims{
		email,
		idreq,
		expiredtime,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(Secretkey)
	if err != nil {
		return err.Error(), 400
	}

	return t, 200
}
func GetValJWT(token *jwt.Token, key string) string {
	// tokenval := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims[key].(string)
}
