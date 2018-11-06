package main

import (
	"github.com/sqshq/piggymetrics/config"
	"github.com/sqshq/piggymetrics/internal/api"
	"github.com/sqshq/piggymetrics/internal/server"
	"github.com/sqshq/piggymetrics/internal/store"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	str := store.New(cfg)
	a := api.Api{Server: srv, Store: str, Config: cfg}
	a.RegisterRoutes()
	srv.Start()
}
