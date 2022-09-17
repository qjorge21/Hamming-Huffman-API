package httpserver

import (
	"github.com/gin-gonic/gin"
	"Hamming-Huffman-API/src/api/internal/controllers/hammingController"
	"net/http"
)

func Routes(router *gin.Engine){
	router.GET("/ping", ping)
	router.POST("/hamming", hamming)
}

func ping(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func hamming(c *gin.Context){
	hammingcontroller.ProtectHamming(c)
}