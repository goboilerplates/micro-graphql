package config

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-rest/middleware"
)

// SetMiddleWares setups middlewares.
func SetMiddleWares(router *gin.Engine) {
	router.Use(middleware.Cors())
	router.Use(middleware.Gzip())
	router.Use(middleware.Static())
}
