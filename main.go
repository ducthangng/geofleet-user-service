package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ducthangng/geofleet/user-service/external"
	"github.com/ducthangng/geofleet/user-service/singleton"
)

var devenv string

func main() {

	// gather the configuration of the service
	cfg := singleton.ReadConfig(devenv)

	// start a global context
	ctx := context.Background()

	// connect to postgre database
	singleton.ConnectPostgre(ctx)

	// connect to redis
	// singleton.ConnectRedis()

	// start routing
	handlers := external.Routing()

	// start the service
	s := &http.Server{
		Handler:           handlers,
		Addr:              cfg.Server.Port,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeout * int(time.Millisecond)),
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout * int(time.Millisecond)),
		WriteTimeout:      time.Duration(cfg.Server.WriteTimeout * int(time.Millisecond)),
		IdleTimeout:       time.Duration(cfg.Server.IdleTimeout * int(time.Millisecond)),
		MaxHeaderBytes:    cfg.Server.MaxHeaderBytes,
	}

	log.Println("start server at port: ", cfg.Server.Port)
	s.ListenAndServe()
}
