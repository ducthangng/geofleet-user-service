package external

import (
	"context"
	"log"

	"github.com/ducthangng/geofleet/user-service/external/middleware"
	"github.com/ducthangng/geofleet/user-service/registry"
	"github.com/gin-gonic/gin"
)

func Routing() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	// middleware

	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.JSONWriterMiddleware)

	userHandler, err := registry.InitializeUserService(context.Background())
	if err != nil {
		// build fail
		log.Println("build failed")
		panic(err)
	}

	// go check for health
	r.GET("/api/healthz")

	// login for all users
	// set cookie with JWT token.
	r.POST("/api/login", userHandler.Login)
	// create new account (does not duplicated phone)
	r.POST("/api/register", userHandler.Register)
	// retrieve the id of the user, require tokens.
	r.GET("/api/user/me")

	return r
}
