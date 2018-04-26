package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-rest/config"
	"github.com/goboilerplates/micro-rest/env"
)

func main() {
	enableProdMode, serverPort := env.LoadVariables()
	if enableProdMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	config.SetMiddleWares(router)
	config.SetAPIPaths(router)

	router.Run(serverPort)
}
