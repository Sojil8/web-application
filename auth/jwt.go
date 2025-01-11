package auth

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey []byte

type User struct {
	UserInfo string `json:"userinfo"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func SecretKey() {
	secret := os.Getenv("Secret_Key")
	if len(secret) == 0 {
		log.Fatal("Faild to get secret key")
	}
	secretKey = []byte(secret)

}

func JwtTokens(c *gin.Context, userinfo string, role string) {
	tokenString, err := CreateToken(userinfo, role)
	if err != nil {
		fmt.Println("Faild to create token")
	}

	session := sessions.Default(c)
	session.Set(role, tokenString)
	session.Save()

	check := session.Get(role)
	fmt.Println("JWT Token: ", check)
}

func CreateToken(userinfo string, role string) (string, error) {
	Claims := User{
		UserInfo: userinfo,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "Error at creating token", nil
	} else {
		return tokenString, nil
	}
}
