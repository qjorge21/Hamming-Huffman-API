package helpers

import (
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
