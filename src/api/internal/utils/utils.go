package utils

import (
	"math"
)

// 2^p >= p + i + 1
func CalculateParityBitsNeeded(i int) int {
	for p := 0; p < i; p += 1 {
		if int(math.Pow(2, float64(p))) < p+i+1 {
			return p
		}
	}

	return 0
}

func InvertirOrdenArreglo(arr []bool) []bool {
	result := make([]bool, len(arr))

	indexResult := len(arr) - 1

	for index, _ := range arr {
		result[index] = arr[indexResult]
		indexResult--
	}

	return result
}

func CalcularValorDecimal(arr []bool) int {
	resultado := 0
	potencia := 0

	for index := len(arr) - 1; index >= 0; index-- {
		if arr[index] {
			resultado += int(math.Pow(2, float64(potencia)))
		}
		potencia++
	}

	return resultado
}

func EsPar(numero int) bool {
	return numero%2 == 0
}
