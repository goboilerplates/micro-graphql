package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-graphql/config"
	"github.com/goboilerplates/micro-graphql/env"
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

	// Serve HTTPS
	// const (
	// 	pemFilePath = "./keys/server.pem"
	// 	keyFilePath = "./keys/server.key"
	// )
	// router.RunTLS(":9002", pemFilePath, keyFilePath)
}
