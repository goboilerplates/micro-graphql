package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

// Cors gets cors middleware.
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		// AllowOrigins:     []string{},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
