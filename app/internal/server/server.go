package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sqshq/piggymetrics-go/app/config"
	"time"
)

type Server struct {
	config *config.Configuration
	Echo   *echo.Echo
}

func New(c *config.Configuration) *Server {

	e := echo.New()
	e.Server.ReadTimeout = time.Duration(c.ReadTimeoutSec) * time.Second
	e.Server.WriteTimeout = time.Duration(c.WriteTimeoutSec) * time.Second
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Recover())
	e.Static("/", "app/assets")

	return &Server{config: c, Echo: e}
}

func (s *Server) Start() {
	s.Echo.Logger.Fatal(s.Echo.Start(":" + s.config.Port))
}
