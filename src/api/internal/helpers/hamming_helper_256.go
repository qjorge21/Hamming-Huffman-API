package helpers

import (
	"Hamming-Huffman-API/src/api/internal/constants"
	"Hamming-Huffman-API/src/api/internal/utils"
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

func CrearArregloDeModulos256(arregloBool []bool, cantModulos int) [][256]bool {
	arregloModulos := make([][256]bool, cantModulos)

	contadorModulo := 0

	for i := 0; i < len(arregloBool); i += 256 {
		indice := 0

		for j := i; j < i+256; j++ {
			arregloModulos[contadorModulo][indice] = arregloBool[j]
			indice++
		}
		contadorModulo++
	}

	return arregloModulos
}

func ChequearErrorModulo256(modulo [256]bool, matriz256 [256][8]bool) int {
	result := make([]bool, constants.TAM_BITS_CONTROL_MODULO_256)

	for columna := 0; columna < constants.TAM_BITS_CONTROL_MODULO_256; columna++ {

		cantidadDeUnos := 0

		for fila := 0; fila < constants.TAM_BITS_TOTALES_MODULO_256; fila++ {

			if matriz256[fila][columna] {

				if modulo[fila] {
					cantidadDeUnos++
				}
			}
		}
		result[columna] = !utils.EsPar(cantidadDeUnos)
	}

	result = utils.InvertirOrdenArreglo(result)
	pos := utils.CalcularValorDecimal(result)

	return pos
}

func CorregirErrorModulo256(arregloModulos [][256]bool, modulo int, pos int) {
	arregloModulos[modulo][pos] = !arregloModulos[modulo][pos]
}

func TransformarArregloModulos256BooleanosToArregloBytes(arreglo [][256]bool) []byte {
	var stringByte = ""
	var contadorBits = 0
	var arregloBytes []byte
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < 256; j++ {
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
