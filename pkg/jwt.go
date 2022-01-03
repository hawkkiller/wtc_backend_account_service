package pkg

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

func NewJWT(expTime time.Duration, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expTime).Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        id,
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
