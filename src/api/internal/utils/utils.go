package utils

import (
	"math"
)

// 2^p >= p + i + 1
func CalculateParityBitsNeeded(i int) int {
	for p:=0; p < i; p += 1 {
		if int(math.Pow(2,float64(p))) < p + i + 1 {
			return p
		}
	}
	
	return 0
}