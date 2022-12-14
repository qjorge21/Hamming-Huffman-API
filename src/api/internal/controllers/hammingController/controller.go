package hammingcontroller

import (
	hammingservice "Hamming-Huffman-API/src/api/internal/services/hammingService"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type respuestaProteger struct {
	TextoProtegerOriginal      string
	TextoProtegerGenerado      string
	TextoProtegerGeneradoBytes []byte
	Tiempo                     float64
}

type respuestaDesproteger struct {
	TextoDesprotegerOriginal      string
	TextoDesprotegerGenerado      string
	TextoDesprotegerGeneradoBytes []byte
	Tiempo                        float64
}

func ProtegerHamming(c *gin.Context) {
	ctx := c.Request.Context()
	parametros := c.Request.URL.Query()

	// TODO: agregar controles de parametros de entrada
	// posible mejora: llevar parametros de url en el body
	if parametros["modulo"][0] == "256" {
		textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.ProtectHamming256(ctx, parametros["file_name"][0], parametros["introducir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaProteger{textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
	if parametros["modulo"][0] == "1024" {
		textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.ProtectHamming1024(ctx, parametros["file_name"][0], parametros["introducir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaProteger{textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
	if parametros["modulo"][0] == "2048" {
		textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.ProtectHamming2048(ctx, parametros["file_name"][0], parametros["introducir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaProteger{textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
	if parametros["modulo"][0] == "4096" {
		textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.ProtectHamming4096(ctx, parametros["file_name"][0], parametros["introducir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaProteger{textoProtegerOriginal, textoProtegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
}

func DesprotegerHamming(c *gin.Context) {
	ctx := c.Request.Context()
	parametros := c.Request.URL.Query()

	if parametros["modulo"][0] == "256" {
		textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.DesprotegerHamming256(ctx, parametros["file_name"][0], parametros["corregir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaDesproteger{textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
	if parametros["modulo"][0] == "1024" {
		textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.DesprotegerHamming1024(ctx, parametros["file_name"][0], parametros["corregir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaDesproteger{textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
	if parametros["modulo"][0] == "2048" {
		textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.DesprotegerHamming2048(ctx, parametros["file_name"][0], parametros["corregir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaDesproteger{textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}
	if parametros["modulo"][0] == "4096" {
		textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo := hammingservice.DesprotegerHamming4096(ctx, parametros["file_name"][0], parametros["corregir_error"][0])
		if tiempo == 0 {
			err := errors.New("error")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		} else {
			var res = respuestaDesproteger{textoDesprotegerOriginal, textoDesprotegerGenerado, textoProtegerGeneradoBytes, tiempo}
			c.JSON(http.StatusOK, res)
		}
	}

}

type BodyProteger struct {
	Archivo []byte `json:"textoProtegerGeneradoBytes"`
}

func SaveFile(c *gin.Context) {
	parametros := c.Request.URL.Query()
	body := BodyProteger{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f, err := os.Create("./internal/texts/" + parametros["file_name"][0])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer f.Close()

	//fmt.Print(body.Archivo)
	//fmt.Print("\n")

	_, err2 := f.Write(body.Archivo)
	if err2 != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	/*
		d2 := []byte{35, 137, 196, 196, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		_, err2 := f.Write(d2)
		_, err2 = f.Write(d2)
		if err2 != nil {
			log.Fatal(err2)
		}
	*/

	c.JSON(http.StatusAccepted, &body)
}

type BodyDesproteger struct {
	TextoDesprotegido []byte `json:"textoDesprotegerGeneradoBytes"`
}

func GuardarTextoDesprotegido(c *gin.Context) {
	parametros := c.Request.URL.Query()
	body := BodyDesproteger{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	f, err := os.Create("./internal/texts/" + parametros["file_name"][0])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defer f.Close()

	_, err2 := f.Write(body.TextoDesprotegido)
	if err2 != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusAccepted, &body)
}
