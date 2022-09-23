package hammingservice

import (
	"Hamming-Huffman-API/src/api/internal/helpers"
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

func ProtectHamming256(ctx context.Context, fileName string, introducir_error string) (string, string, []byte, float64) {
	fmt.Print("Nombre del Achivo a Proteger: ")
	fmt.Print(fileName)
	fmt.Print("\n")

	start := time.Now()

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 247, 8
	const MODULO = CANTIDAD_BITS_INFO + CANTIDAD_BITS_CONTROL + 1

	var textoProtegerGeneradoBytes []byte
	fileAsBytes, err := ioutil.ReadFile("./internal/texts/" + fileName)
	if err != nil {
		return "", "", textoProtegerGeneradoBytes, 0
	}

	//fmt.Printf("File contents: %s", fileAsBytes)
	//fmt.Print("\n")

	// Imprimir el archivo por bytes (en decimal)
	//fmt.Print(fileAsBytes)
	//fmt.Print("\n")

	var length = len(fileAsBytes)
	fmt.Print("Cantidad de Bytes: ")
	fmt.Print(length)
	fmt.Print("\n")

	fmt.Print("Cantidad de Bits de Información: ")
	var cantidadBitsInformacion = length * 8
	fmt.Print(cantidadBitsInformacion)
	fmt.Print("\n")

	fmt.Print("Cantidad Módulos ")
	fmt.Print(MODULO)
	fmt.Print(": ")
	var cantidadModulos = int(math.Ceil(float64(cantidadBitsInformacion) / CANTIDAD_BITS_INFO))
	fmt.Print(cantidadModulos)
	fmt.Print("\n")

	arrayModules := make([][MODULO]bool, cantidadModulos)

	arrayBitsInformacion := make([]bool, cantidadBitsInformacion)
	indexMatriz := 0
	byteObtenido := ""

	for _, n := range fileAsBytes {
		byteObtenido = fmt.Sprintf("%08b", n)
		//fmt.Print(byteObtenido)
		//fmt.Print("\n")

		for _, value := range byteObtenido {
			if string(value) == "1" {
				arrayBitsInformacion[indexMatriz] = true
			} else {
				arrayBitsInformacion[indexMatriz] = false
			}
			indexMatriz = indexMatriz + 1
		}
	}

	/*
		fmt.Print("Arreglo Bits de Información\n")
		var contador = 0
		for row := 0; row < cantidadBitsInformacion; row++ {
			if arrayBitsInformacion[row] {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}

			if (contador + 1) == 8 {
				fmt.Print(" ")
				contador = 0
			} else {
				contador++
			}
		}
		fmt.Print("\n")
	*/

	//fmt.Println(arrayModules)

	// Este 'for' llena el modulo de 256 saltando los lugares donde van los bits de control
	var contadorPosicionModulo = 1
	var contadorModulo = 0
	for index := 0; index < cantidadBitsInformacion; index++ {
		var potenciaDos int
		potenciaDos = (contadorPosicionModulo & (contadorPosicionModulo - 1))
		for potenciaDos == 0 {
			contadorPosicionModulo = contadorPosicionModulo + 1
			potenciaDos = (contadorPosicionModulo & (contadorPosicionModulo - 1))
		}
		arrayModules[contadorModulo][contadorPosicionModulo] = arrayBitsInformacion[index]

		contadorPosicionModulo = contadorPosicionModulo + 1

		if contadorPosicionModulo == MODULO {
			contadorModulo = contadorModulo + 1
			/*
				fmt.Print("Se completo el modulo: ")
				fmt.Print(contadorModulo)
				fmt.Print("\n")
				fmt.Print("Se introdujo el bit de info: ")
				fmt.Print(index + 1)
				fmt.Print("\n")
			*/
			contadorPosicionModulo = 1
		}
	}
	//fmt.Println(arrayModules)
	/*
		fmt.Print("Arreglo de Modulos Sin Calcular Bits de Control:\n")
		fmt.Print("[")
		for modulo := 0; modulo < cantidadModulos; modulo++ {
			fmt.Print("[")
			for posicionBit := 0; posicionBit < MODULO; posicionBit++ {
				var potenciaDos int
				potenciaDos = (posicionBit & (posicionBit - 1))
				if potenciaDos != 0 {
					if arrayModules[modulo][posicionBit] {
						fmt.Print("1")
					} else {
						fmt.Print("0")
					}
				} else {
					if arrayModules[modulo][posicionBit] {
						fmt.Print("(1)")
					} else {
						fmt.Print("(0)")
					}
				}
			}
			fmt.Print("]")
		}
		fmt.Print("]")
		fmt.Print("\n")
		fmt.Print("\n")
	*/
	// Este 'for' calcula los bits de control en las posiciones potencias de dos
	var matriz256 = helpers.GenerarMatriz256()
	for modulo := 0; modulo < cantidadModulos; modulo++ {
		var potenciaDos = 1
		var columna = 0
		for potenciaDos < MODULO {
			var paridad = false
			for fila := 0; fila < MODULO; fila++ {
				if matriz256[fila][columna] {
					if fila != potenciaDos {
						if arrayModules[modulo][fila] {
							if paridad {
								paridad = false
							} else {
								paridad = true
							}
						}
					}
				}
			}
			arrayModules[modulo][potenciaDos] = paridad
			columna = columna + 1
			potenciaDos = potenciaDos * 2
		}
	}

	/*
		fmt.Print("Arreglo de Modulos con Bits de Control Calculados:\n")
		fmt.Print("[")
		for modulo := 0; modulo < cantidadModulos; modulo++ {
			fmt.Print("[")
			for posicionBit := 0; posicionBit < MODULO; posicionBit++ {
				var potenciaDos int
				potenciaDos = (posicionBit & (posicionBit - 1))
				if potenciaDos != 0 {
					if arrayModules[modulo][posicionBit] {
						fmt.Print("1")
					} else {
						fmt.Print("0")
					}
				} else {
					if arrayModules[modulo][posicionBit] {
						fmt.Print("(1)")
					} else {
						fmt.Print("(0)")
					}
				}
			}
			fmt.Print("]")
		}
		fmt.Print("]")
		fmt.Print("\n")
		fmt.Print("\n")
	*/

	var textoProtegido string
	for k := 0; k < len(arrayModules); k++ {
		if len(textoProtegido) == 0 {
			textoProtegido = helpers.TransformarArreglo256BooleanosToString(arrayModules[k])
		} else {
			textoProtegido = textoProtegido + helpers.TransformarArreglo256BooleanosToString(arrayModules[k])
		}
	}

	textoProtegerGeneradoBytes = helpers.TransformarArregloModulos256BooleanosToArregloBytes(arrayModules)
	//fmt.Print(textoProtegerGeneradoBytes)
	//fmt.Print("\n")

	if introducir_error == "true" {
		fmt.Print("Introducir error...")
		fmt.Print("\n")
		// Calcular modulo random y posicion del modulo random -> Cambiar ese bit
	}

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoProtegido, textoProtegerGeneradoBytes, elapsed
}

func BytesInFile(fileName string) {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)

	// Call Split to specify that we want to Scan each individual byte.
	scanner.Split(bufio.ScanBytes)

	// Use For-loop.
	for scanner.Scan() {
		// Get Bytes and display the byte.
		//b := scanner.Bytes()
		//fmt.Printf("%v = %v = %v\n", b, b[0], string(b))
	}
}

func DesprotegerHamming256(ctx context.Context, fileName string, corregir_error string) (string, string, []byte, float64) {
	// Iniciar el timer
	start := time.Now()

	// Leer el archivo
	fmt.Print("Nombre del Achivo a Desproteger: ")
	fmt.Print(fileName)
	fmt.Print("\n")
	var textoDesprotegerGeneradoBytes []byte
	fileAsBytes, err := ioutil.ReadFile("./internal/texts/" + fileName)
	if err != nil {
		return "", "", textoDesprotegerGeneradoBytes, 0
	}

	//fmt.Printf("File contents: %s", fileAsBytes)
	//fmt.Print("\n")

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 247, 8
	const MODULO = CANTIDAD_BITS_INFO + CANTIDAD_BITS_CONTROL + 1
	const CANTIDAD_BYTES_MODULO = MODULO / 8

	//fmt.Print(fileAsBytes)
	//fmt.Print("\n")

	var length = len(fileAsBytes)
	fmt.Print("Cantidad de Bytes: ")
	fmt.Print(length)
	fmt.Print("\n")

	fmt.Print("Cantidad de Módulos ")
	fmt.Print(MODULO)
	fmt.Print(": ")
	var cantidadModulos = int(math.Ceil(float64(length) / CANTIDAD_BYTES_MODULO))
	fmt.Print(cantidadModulos)
	fmt.Print("\n")

	fmt.Print("Cantidad de Bits de Información: ")
	var cantidadBitsInformacion = cantidadModulos * 247
	fmt.Print(cantidadBitsInformacion)
	fmt.Print("\n")

	byteObtenido := ""
	arrayBitsInformacion := make([]bool, cantidadBitsInformacion)
	indexModulo256 := 0
	indexInfo := 0
	controlBytesModulo := 0

	for _, n := range fileAsBytes {
		byteObtenido = fmt.Sprintf("%08b", n)
		for _, value := range byteObtenido {
			// Si es un bit de control lo tenemos que ignorar
			var potenciaDos int
			potenciaDos = (indexModulo256 & (indexModulo256 - 1))
			if potenciaDos == 0 || indexModulo256 == 0 {
				/*
					fmt.Print("Saltamos la pos : ")
					fmt.Print(indexModulo256)
					fmt.Print("\n")
				*/
				indexModulo256 = indexModulo256 + 1
				continue
			}

			if string(value) == "1" {
				arrayBitsInformacion[indexInfo] = true
			} else {
				arrayBitsInformacion[indexInfo] = false
			}
			indexInfo = indexInfo + 1
			indexModulo256 = indexModulo256 + 1
		}

		controlBytesModulo = controlBytesModulo + 1

		// Si ya se consumió un modulo entero hay que reiniciar las variables
		if controlBytesModulo == CANTIDAD_BYTES_MODULO {
			indexModulo256 = 0
			controlBytesModulo = 0
		}
	}

	var textoDesprotegido string
	textoDesprotegido = helpers.TransformarArregloBooleanosToString(arrayBitsInformacion)
	textoDesprotegerGeneradoBytes = helpers.TransformarArregloBooleanosToArregloBytes(arrayBitsInformacion)
	/*
		fmt.Print("Arreglo Bits de Información\n")
		var contador = 0
		for row := 0; row < cantidadBitsInformacion; row++ {
			if arrayBitsInformacion[row] {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}

			if (contador + 1) == 8 {
				fmt.Print(" ")
				contador = 0
			} else {
				contador++
			}
		}
		fmt.Print("\n")
	*/

	if corregir_error == "true" {
		// Buscar y corregir error
	}

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoDesprotegido, textoDesprotegerGeneradoBytes, elapsed
}
