package crypto

import (
	"fmt"
	"log"
	"time"

	"github.com/bingoohuang/go-rest-template/internal/pkg/conf"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), plainPwd) == nil
}

func CreateToken(username string) (string, error) {
	c := conf.GetConf()

	var err error
	// Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	token, err := at.SignedString([]byte(c.Server.Secret)) // SECRET
	if err != nil {
		return "token creation error", err
	}
	return token, nil
}

func ValidateToken(tokenString string) bool {
	c := conf.GetConf()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(c.Server.Secret), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}
