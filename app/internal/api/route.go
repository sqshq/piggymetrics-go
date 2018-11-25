package api

import (
	"github.com/labstack/echo/middleware"
	"github.com/sqshq/piggymetrics-go/app/config"
	"github.com/sqshq/piggymetrics-go/app/internal/server"
	"github.com/sqshq/piggymetrics-go/app/internal/store"
)

type Api struct {
	Server *server.Server
	Store  *store.Store
	Config *config.Configuration
}

func (a *Api) RegisterRoutes() {
	a.Server.Echo.GET("/healthcheck", a.Healthcheck)
	a.Server.Echo.POST("/uaa/oauth/token", a.CreateToken)

	g := a.Server.Echo.Group("/accounts/current")
	g.Use(middleware.JWT([]byte(a.Config.JwtSecret)))
	g.GET("", a.GetCurrentAccount)
	g.PUT("", a.SaveCurrentAccount)

	a.Server.Echo.POST("/accounts/", a.CreateNewAccount)
	a.Server.Echo.GET("/accounts/demo", a.GetDemoAccount)

	a.Server.Echo.PUT("/notifications/recipients/current", a.SubscribeForNotifications)
}
