package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sqshq/piggymetrics/internal/model/account"
	"github.com/sqshq/piggymetrics/internal/model/user"
	"net/http"
)

func (a *Api) Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (a *Api) GetDemoAccount(c echo.Context) error {
	return c.JSONBlob(http.StatusOK, []byte(a.Config.DemoAccountDump))
}

func (a *Api) CreateNewAccount(c echo.Context) error {
	u := new(user.User)

	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Can't deserialize a user")
	}

	_, err := user.Create(a.Store, u)
	if err != nil {
		log.Error("Failed to create a user ", err)
		return c.JSON(http.StatusInternalServerError, "Failed to create a user, please try again later")
	}

	acc, err := account.Create(a.Store, u)

	if err != nil {
		log.Error("Failed to create an account ", err)
		return c.JSON(http.StatusInternalServerError, "Failed to create an account, please try again later")
	}

	return c.JSON(http.StatusOK, acc)
}

func (a *Api) GetCurrentAccount(c echo.Context) error {

	token := DecodeToken(c)
	if token.expired {
		return c.JSON(http.StatusForbidden, "Token expired")
	}

	acc := account.FindByName(a.Store, token.username)
	return c.JSON(http.StatusOK, acc)
}

func (a *Api) SaveCurrentAccount(c echo.Context) error {

	token := DecodeToken(c)
	if token.expired {
		return c.JSON(http.StatusForbidden, "Token expired")
	}

	acc := new(account.Account)
	if err := c.Bind(acc); err != nil {
		return c.JSON(http.StatusBadRequest, "Can't deserialize an account: ")
	}

	err := account.Update(a.Store, acc)

	if err != nil {
		log.Error("Failed to create a user ", err)
		return c.JSON(http.StatusInternalServerError, "Failed to update an account, please try again later")
	}

	return c.JSON(http.StatusOK, "OK")
}

func (a *Api) CreateToken(c echo.Context) error {

	u := new(user.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, "Can't deserialize a user")
	}

	authenticated := user.Authenticate(a.Store, u)
	if !authenticated {
		return c.JSON(http.StatusForbidden, "Forbidden")
	}

	t, err := CreateToken(a.Config, u)
	if err != nil {
		log.Error("Failed to create a token ", err)
		return c.JSON(http.StatusInternalServerError, "Can't create a token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": t,
	})
}

func (a *Api) SubscribeForNotifications(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Notifications functionality is currently unavailable")
}
