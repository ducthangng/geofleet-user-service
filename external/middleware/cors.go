package middleware

import (
	"sync"
	"time"

	// "github.com/!tech-by-!g!l/!telegram!bot/singleton"

	"github.com/ducthangng/geofleet/user-service/singleton"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	corsOnce       sync.Once
	corsMiddleware gin.HandlerFunc
)

// CorsMiddleware return the middeware instance
func CorsMiddleware() gin.HandlerFunc {

	cfg := singleton.GetConfig().Server

	corsOnce.Do(func() {
		corsMiddleware = cors.New(cors.Config{
			AllowOrigins: []string{cfg.HTTPDomain},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders: []string{
				"Content-Length",
				"Content-Type",
				"Access-Control-Allow-Headers",
				"Access-Control-Allow-Origin",
				"Origin",
				"Accept-Encoding",
				"X-CSRF-Token",
				"Authorization",
				"*",
			},
			ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		})
	})
	return corsMiddleware
}
