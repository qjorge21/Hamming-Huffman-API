package helpers

import (
	"fmt"
	"math"
)

func GenerarMatriz256() [256][8]bool {
	var matriz [256][8]bool
	var salto = 1
	for columna := 0; columna < 8; columna++ {
		var fila = int(math.Pow(2, float64(columna)))
		var control_salto, saltados = 0, 0
		for ; fila < 256; fila++ {
			if salto == control_salto {
				if salto == saltados {
					matriz[fila][columna] = true
					control_salto, saltados = 1, 0
				} else {
					saltados++
				}
			} else {
				matriz[fila][columna] = true
				control_salto++
			}
		}
		salto = salto * 2
	}
	return matriz
}

func GenerarMatriz1024() [1024][10]bool {
	var matriz [1024][10]bool
	var salto = 1
	for columna := 0; columna < 10; columna++ {
		var fila = int(math.Pow(2, float64(columna)))
		var control_salto, saltados = 0, 0
		for ; fila < 1024; fila++ {
			if salto == control_salto {
				if salto == saltados {
					matriz[fila][columna] = true
					control_salto, saltados = 1, 0
				} else {
					saltados++
				}
			} else {
				matriz[fila][columna] = true
				control_salto++
			}
		}
		salto = salto * 2
	}
	return matriz
}

func GenerarMatriz2048() [2048][11]bool {
	var matriz [2048][11]bool
	var salto = 1
	for columna := 0; columna < 11; columna++ {
		var fila = int(math.Pow(2, float64(columna)))
		var control_salto, saltados = 0, 0
		for ; fila < 2048; fila++ {
			if salto == control_salto {
				if salto == saltados {
					matriz[fila][columna] = true
					control_salto, saltados = 1, 0
				} else {
					saltados++
				}
			} else {
				matriz[fila][columna] = true
				control_salto++
			}
		}
		salto = salto * 2
	}
	return matriz
}

func GenerarMatriz4096() [4096][12]bool {
	var matriz [4096][12]bool
	var salto = 1
	for columna := 0; columna < 12; columna++ {
		var fila = int(math.Pow(2, float64(columna)))
		var control_salto, saltados = 0, 0
		for ; fila < 4096; fila++ {
			if salto == control_salto {
				if salto == saltados {
					matriz[fila][columna] = true
					control_salto, saltados = 1, 0
				} else {
					saltados++
				}
			} else {
				matriz[fila][columna] = true
				control_salto++
			}
		}
		salto = salto * 2
	}
	return matriz
}

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

func TransformarArregloModulos1024BooleanosToArregloBytes(arreglo [][1024]bool) []byte {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < 1024; j++ {
			if arreglo[i][j] {
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
	}
	return arregloBytes
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

func TransformarArregloModulos2048BooleanosToArregloBytes(arreglo [][2048]bool) []byte {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < 2048; j++ {
			if arreglo[i][j] {
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
	}
	return arregloBytes
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

func TransformarArregloModulos4096BooleanosToArregloBytes(arreglo [][4096]bool) []byte {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < 4096; j++ {
			if arreglo[i][j] {
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
	}
	return arregloBytes
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
