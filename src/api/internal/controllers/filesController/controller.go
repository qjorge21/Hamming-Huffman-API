package filescontroller

import (
	"Hamming-Huffman-API/src/api/internal/services/filesService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFile(c *gin.Context){
	ctx := c.Request.Context()

	fileString, err := filesservice.ReadFile(ctx)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, fileString)
	}

	c.JSON(http.StatusOK, fileString)
}