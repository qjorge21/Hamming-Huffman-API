package httpserver

import (
	"github.com/gin-gonic/gin"
	"Hamming-Huffman-API/src/api/internal/controllers/filesController"
	"net/http"
)

func Routes(router *gin.Engine){
	router.GET("/ping", ping)
	router.POST("/readFile", readFile)
}

func ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func readFile(c *gin.Context){
	filescontroller.ReadFile(c)
}