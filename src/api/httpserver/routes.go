package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine){
	router.GET("/ping", ping)
}

func ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}