package main

import (
	"github.com/sqshq/piggymetrics-go/app/config"
	"github.com/sqshq/piggymetrics-go/app/internal/api"
	"github.com/sqshq/piggymetrics-go/app/internal/server"
	"github.com/sqshq/piggymetrics-go/app/internal/store"
)

func main() {
	cfg := config.Load()
	srv := server.New(cfg)
	str := store.New(cfg)
	a := api.Api{Server: srv, Store: str, Config: cfg}
	a.RegisterRoutes()
	srv.Start()
}
