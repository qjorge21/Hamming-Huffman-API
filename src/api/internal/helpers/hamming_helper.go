package helpers

import (
	"fmt"
	"math"
)

func TransformarStringByteToDecimal(stringByte string) int {
	var decimalCorrespondiente = 0
	var potencia = 7
	for j := 0; j < 8; j++ {
		if string(stringByte[j]) == "1" {
			decimalCorrespondiente = decimalCorrespondiente + int(math.Pow(2, float64(potencia)))
		}
		potencia = potencia - 1
	}
	return decimalCorrespondiente
}

func TransformarArregloBooleanosToString(arreglo []bool) string {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] {
			stringByte = stringByte + "1"
		} else {
			stringByte = stringByte + "0"
		}
		contadorBits = contadorBits + 1
		if contadorBits == 8 {
			arregloBytes = append(arregloBytes, byte(TransformarStringByteToDecimal(stringByte)))
			stringByte = ""
			contadorBits = 0
		}
	}
	return string(arregloBytes)
}

func TransformarArreglo256BooleanosToString(arreglo [256]bool) string {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] {
			stringByte = stringByte + "1"
		} else {
			stringByte = stringByte + "0"
		}
		contadorBits = contadorBits + 1
		if contadorBits == 8 {
			arregloBytes = append(arregloBytes, byte(TransformarStringByteToDecimal(stringByte)))
			stringByte = ""
			contadorBits = 0
		}
	}
	return string(arregloBytes)
}

func TransformarArregloBooleanosToArregloBytes(arreglo []bool) []byte {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] {
			stringByte = stringByte + "1"
		} else {
			stringByte = stringByte + "0"
		}
		contadorBits = contadorBits + 1
		if contadorBits == 8 {
			arregloBytes = append(arregloBytes, byte(TransformarStringByteToDecimal(stringByte)))
			stringByte = ""
			contadorBits = 0
		}
	}
	return arregloBytes
}

func TransformarArreglo1024BooleanosToString(arreglo [1024]bool) string {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] {
			stringByte = stringByte + "1"
		} else {
			stringByte = stringByte + "0"
		}
		contadorBits = contadorBits + 1
		if contadorBits == 8 {
			arregloBytes = append(arregloBytes, byte(TransformarStringByteToDecimal(stringByte)))
			stringByte = ""
			contadorBits = 0
		}
	}
	return string(arregloBytes)
}

func TransformarArreglo2048BooleanosToString(arreglo [2048]bool) string {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] {
			stringByte = stringByte + "1"
		} else {
			stringByte = stringByte + "0"
		}
		contadorBits = contadorBits + 1
		if contadorBits == 8 {
			arregloBytes = append(arregloBytes, byte(TransformarStringByteToDecimal(stringByte)))
			stringByte = ""
			contadorBits = 0
		}
	}
	return string(arregloBytes)
}

func TransformarArreglo4096BooleanosToString(arreglo [4096]bool) string {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] {
			stringByte = stringByte + "1"
		} else {
			stringByte = stringByte + "0"
		}
		contadorBits = contadorBits + 1
		if contadorBits == 8 {
			arregloBytes = append(arregloBytes, byte(TransformarStringByteToDecimal(stringByte)))
			stringByte = ""
			contadorBits = 0
		}
	}
	return string(arregloBytes)
}

func TransformarArregloBytesToArregloBool(arregloBytes []byte) []bool {
	byteString := ""
	indice := 0
	arregloBool := make([]bool, len(arregloBytes)*8)

	for _, n := range arregloBytes {
		byteString = fmt.Sprintf("%08b", n)

		for _, bit := range byteString {
			if string(bit) == "1" {
				arregloBool[indice] = true
			} else {
				arregloBool[indice] = false
			}
			indice++
		}
	}

	return arregloBool
}

func CalcularCantidadModulos(archivoBool []bool, tamModulo int) int {
	return len(archivoBool) / tamModulo
}
