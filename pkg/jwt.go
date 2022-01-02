package pkg

import (
	"github.com/golang-jwt/jwt"
	"main/internal/data/model"
	"os"
	"time"
)

func NewJWT(args ...model.MapArg) (string, error) {
	claims := jwt.MapClaims{}
	for i := 0; i < len(args); i++ {
		claims[args[i].Key] = args[i].Value
	}
	claims["time"] = time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(os.Getenv("SECRET")))
}
