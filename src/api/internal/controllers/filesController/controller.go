package filescontroller

import (
	"github.com/gin-gonic/gin"
	"Hamming-Huffman-API/src/api/internal/services/filesService"
)

func ReadFile(c *gin.Context){
	ctx := c.Request.Context()
	filesservice.ReadFile(ctx)
}