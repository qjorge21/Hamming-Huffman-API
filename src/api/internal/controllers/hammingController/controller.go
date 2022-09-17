package hammingcontroller

import (
	"Hamming-Huffman-API/src/api/internal/services/hammingService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectHamming(c *gin.Context){
	ctx := c.Request.Context()
	hammingservice.ProtectHamming(ctx)
	c.JSON(http.StatusOK, "OK")
}