package httpserver

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	SetupRouter().Run(":8080")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	Routes(router)
	return router
}
