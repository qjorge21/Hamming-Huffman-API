package main

import (
	"Hamming-Huffman-API/src/api/httpserver"
	"Hamming-Huffman-API/src/api/internal/helpers"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var matriz256 = helpers.GenerarMatriz256()
	var matriz1024 = helpers.GenerarMatriz1024()
	var matriz2048 = helpers.GenerarMatriz2048()
	var matriz4096 = helpers.GenerarMatriz4096()

	// Mostrar matriz por consola
	/*
		var modulo, paridad = 256, 8
		for row := 0; row < modulo; row++ {
			for column := 0; column < paridad; column++ {
				if matriz256[row][column] {
					fmt.Print(1, " ")
				} else {
					fmt.Print(0, " ")
				}
			}
			fmt.Print("\n")
		}
	*/
	if matriz256[255][7] || matriz1024[1023][9] || matriz2048[2047][10] || matriz4096[4095][11] {
		// Este if sólo está para que no me de error porque las variables no están utilizadas...
	}

	elapsed := time.Since(start)
	fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)

	httpserver.Start()
}
