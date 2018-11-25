package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/sqshq/piggymetrics-go/app/config"
	"github.com/sqshq/piggymetrics-go/app/internal/model/user"
	"time"
)

type Token struct {
	username string
	expired  bool
}

func CreateToken(c *config.Configuration, u *user.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte(c.JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func DecodeToken(c echo.Context) Token {
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(jwt.MapClaims)
	un := claims["username"].(string)
	exp := time.Unix(int64(claims["exp"].(float64)), 0)
	return Token{username: un, expired: time.Now().After(exp)}
}
