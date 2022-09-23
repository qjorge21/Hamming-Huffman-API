package httpserver

import (
	hammingcontroller "Hamming-Huffman-API/src/api/internal/controllers/hammingController"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Routes(router *gin.Engine) {
	router.Use(CORSMiddleware())
	router.GET("/ping", ping)
	router.GET("/protegerHamming", protegerHamming)
	router.GET("/desprotegerHamming", desprotegerHamming)
	router.POST("/saveFile", saveFile)
	router.POST("/guardarTextoDesprotegido", guardarTextoDesprotegido)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func protegerHamming(c *gin.Context) {
	hammingcontroller.ProtegerHamming(c)
}

func desprotegerHamming(c *gin.Context) {
	hammingcontroller.DesprotegerHamming(c)
}

func saveFile(c *gin.Context) {
	hammingcontroller.SaveFile(c)
}

func guardarTextoDesprotegido(c *gin.Context) {
	hammingcontroller.GuardarTextoDesprotegido(c)
}
