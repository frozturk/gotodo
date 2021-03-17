package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/frozturk/gologin/redis"
)

func CreateToken(userid uint) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ATSECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
