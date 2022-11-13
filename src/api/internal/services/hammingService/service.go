package hammingservice

import (
	"Hamming-Huffman-API/src/api/internal/constants"
	"Hamming-Huffman-API/src/api/internal/helpers"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"time"
)

func ProtectHamming256(ctx context.Context, fileName string, introducir_error string) (string, string, []byte, float64) {
	fmt.Print("Nombre del Achivo a Proteger: ")
	fmt.Print(fileName)
	fmt.Print("\n")

	// Inicia el contador de tiempo
	start := time.Now()

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 247, 8
	const MODULO = CANTIDAD_BITS_INFO + CANTIDAD_BITS_CONTROL + 1

	// Lee el archivo en bytes
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

	// Calcula la cantidad de bits de informacion del archivo de entrada
	fmt.Print("Cantidad de Bits de Información: ")
	var cantidadBitsInformacion = length * 8
	fmt.Print(cantidadBitsInformacion)
	fmt.Print("\n")

	// Calcula la cantidad de modulos de 256 bits necesarios para todos los bits de informacion
	fmt.Print("Cantidad Módulos ")
	fmt.Print(MODULO)
	fmt.Print(": ")
	var cantidadModulos = int(math.Ceil(float64(cantidadBitsInformacion) / CANTIDAD_BITS_INFO))
	fmt.Print(cantidadModulos)
	fmt.Print("\n")

	// Crea un slice de tamaño "cantidadModulos" donde en cada posicion tiene un arreglo de tamaño 256
	arrayModules := make([][MODULO]bool, cantidadModulos)

	// Crea un slice de booleanos de tamaño igual a la cantidad de bits de informacion del archivo de entrada
	arrayBitsInformacion := make([]bool, cantidadBitsInformacion)
	indexMatriz := 0
	byteObtenido := ""

	// Itera sobre los bytes del archivo de entrada y
	// llena el slice de booleanos con 'true' o 'false' segun el bit de informacion sea '1' o '0'
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

	// Llena el slice de modulos de 256 bits con los bits de informacion saltando el primer lugar y
	// aquellos lugares donde se ubican bits de control (posiciones potencias de dos)
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

	// Calcula y setea los bits de control en las posiciones potencias de dos en todos los modulos de 256
	var matriz256 = helpers.GenerarMatriz256()
	for modulo := 0; modulo < cantidadModulos; modulo++ {
		var potenciaDos = 1
		var columna = 0
		for potenciaDos < MODULO { // Mientras la potencia de dos sea menor a 256
			var paridad = false
			for fila := 0; fila < MODULO; fila++ {
				// recorre cada columna matriz 256
				// solo hace algo si el valor es un 1
				if matriz256[fila][columna] { // Recorre la columna que indica qué bits de información deben tenerse en cuenta para calcular la paridad del bit de control
					if fila != potenciaDos {
						if arrayModules[modulo][fila] { // Si hay 'true' en la posicion se actualiza el calculo de paridad
							if paridad {
								paridad = false
							} else {
								paridad = true
							}
						}
					}
				}
			}
			arrayModules[modulo][potenciaDos] = paridad // Setea el bit de control en su posicion
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

	// Si el usuario indicó que hay que introducir errores
	if introducir_error == "true" {
		fmt.Print("Introducir error...")
		fmt.Print("\n")
		s1 := rand.NewSource(time.Now().UnixNano()) // Setea la semilla
		for modulo := 0; modulo < cantidadModulos; modulo++ {
			r1 := rand.New(s1)
			introducirError := r1.Intn(2) // Calcula un random n (donde 0 <= n < 2) que indica si el modulo va a tener error o no
			if introducirError == 1 {
				posicionError := r1.Intn(255) + 1 // Calcula un random n (donde 1 <= n < 256) que indica en qué posicion del modulo va a ser el error
				fmt.Printf("Se introdujo error en el modulo %d y en la posicion %d \n", modulo, posicionError)
				//str := fmt.Sprintf("Mod: %d, Pos: %d, Bit: %d\n", modulo, posicionError, arrayModules[modulo][posicionError])
				//fmt.Print(str)
				arrayModules[modulo][posicionError] = !arrayModules[modulo][posicionError] // Setea el error
				//str = fmt.Sprintf("Bit: %d\n", arrayModules[modulo][posicionError])
				//fmt.Print(str)
			}
		}
	}

	// Obtiene el texto protegido en string transformando modulos de 256 booleanos a strings y concatenandolos
	var textoProtegido string
	for k := 0; k < len(arrayModules); k++ {
		if len(textoProtegido) == 0 {
			textoProtegido = helpers.TransformarArreglo256BooleanosToString(arrayModules[k])
		} else {
			textoProtegido = textoProtegido + helpers.TransformarArreglo256BooleanosToString(arrayModules[k])
		}
	}

	// Obtiene el texto protegido en bytes
	textoProtegerGeneradoBytes = helpers.TransformarArregloModulos256BooleanosToArregloBytes(arrayModules)
	//fmt.Print(textoProtegerGeneradoBytes)
	//fmt.Print("\n")

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoProtegido, textoProtegerGeneradoBytes, elapsed
}

func ProtectHamming1024(ctx context.Context, fileName string, introducir_error string) (string, string, []byte, float64) {
	fmt.Print("Nombre del Achivo a Proteger: ")
	fmt.Print(fileName)
	fmt.Print("\n")

	start := time.Now()

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 1013, 10
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
	var matriz1024 = helpers.GenerarMatriz1024()
	for modulo := 0; modulo < cantidadModulos; modulo++ {
		var potenciaDos = 1
		var columna = 0
		for potenciaDos < MODULO {
			var paridad = false
			for fila := 0; fila < MODULO; fila++ {
				if matriz1024[fila][columna] {
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

	if introducir_error == "true" {
		fmt.Print("Introducir error...")
		fmt.Print("\n")
		s1 := rand.NewSource(time.Now().UnixNano()) // Setea la semilla
		for modulo := 0; modulo < cantidadModulos; modulo++ {
			r1 := rand.New(s1)
			introducirError := r1.Intn(2) // Calcula un random n, donde 0 <= n < 2
			if introducirError == 1 {
				posicionError := r1.Intn(1023) + 1 // Calcula un random n, donde 1 <= n < 1024
				//str := fmt.Sprintf("Mod: %d, Pos: %d, Bit: %d\n", modulo, posicionError, arrayModules[modulo][posicionError])
				//fmt.Print(str)
				fmt.Printf("Se introdujo error en el modulo %d y en la posicion %d \n", modulo, posicionError)
				arrayModules[modulo][posicionError] = !arrayModules[modulo][posicionError]
				//str = fmt.Sprintf("Bit: %d\n", arrayModules[modulo][posicionError])
				//fmt.Print(str)
			}
		}
	}

	var textoProtegido string
	for k := 0; k < len(arrayModules); k++ {
		if len(textoProtegido) == 0 {
			textoProtegido = helpers.TransformarArreglo1024BooleanosToString(arrayModules[k])
		} else {
			textoProtegido = textoProtegido + helpers.TransformarArreglo1024BooleanosToString(arrayModules[k])
		}
	}

	textoProtegerGeneradoBytes = helpers.TransformarArregloModulos1024BooleanosToArregloBytes(arrayModules)
	//fmt.Print(textoProtegerGeneradoBytes)
	//fmt.Print("\n")

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoProtegido, textoProtegerGeneradoBytes, elapsed
}

func ProtectHamming2048(ctx context.Context, fileName string, introducir_error string) (string, string, []byte, float64) {
	fmt.Print("Nombre del Achivo a Proteger: ")
	fmt.Print(fileName)
	fmt.Print("\n")

	start := time.Now()

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 2036, 11
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

	// Este 'for' llena el modulo de 2048 saltando los lugares donde van los bits de control
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
	var matriz2048 = helpers.GenerarMatriz2048()
	for modulo := 0; modulo < cantidadModulos; modulo++ {
		var potenciaDos = 1
		var columna = 0
		for potenciaDos < MODULO {
			var paridad = false
			for fila := 0; fila < MODULO; fila++ {
				if matriz2048[fila][columna] {
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

	if introducir_error == "true" {
		fmt.Print("Introducir error...")
		fmt.Print("\n")
		s1 := rand.NewSource(time.Now().UnixNano()) // Setea la semilla
		for modulo := 0; modulo < cantidadModulos; modulo++ {
			r1 := rand.New(s1)
			introducirError := r1.Intn(2) // Calcula un random n, donde 0 <= n < 2
			if introducirError == 1 {
				posicionError := r1.Intn(2047) + 1 // Calcula un random n, donde 1 <= n < 2048
				//str := fmt.Sprintf("Mod: %d, Pos: %d, Bit: %d\n", modulo, posicionError, arrayModules[modulo][posicionError])
				//fmt.Print(str)
				arrayModules[modulo][posicionError] = !arrayModules[modulo][posicionError]
				//str = fmt.Sprintf("Bit: %d\n", arrayModules[modulo][posicionError])
				//fmt.Print(str)
			}
		}
	}

	var textoProtegido string
	for k := 0; k < len(arrayModules); k++ {
		if len(textoProtegido) == 0 {
			textoProtegido = helpers.TransformarArreglo2048BooleanosToString(arrayModules[k])
		} else {
			textoProtegido = textoProtegido + helpers.TransformarArreglo2048BooleanosToString(arrayModules[k])
		}
	}

	textoProtegerGeneradoBytes = helpers.TransformarArregloModulos2048BooleanosToArregloBytes(arrayModules)
	//fmt.Print(textoProtegerGeneradoBytes)
	//fmt.Print("\n")

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoProtegido, textoProtegerGeneradoBytes, elapsed
}

func ProtectHamming4096(ctx context.Context, fileName string, introducir_error string) (string, string, []byte, float64) {
	fmt.Print("Nombre del Achivo a Proteger: ")
	fmt.Print(fileName)
	fmt.Print("\n")

	start := time.Now()

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 4083, 12
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

	// Este 'for' llena el modulo de 4096 saltando los lugares donde van los bits de control
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
	var matriz4096 = helpers.GenerarMatriz4096()
	for modulo := 0; modulo < cantidadModulos; modulo++ {
		var potenciaDos = 1
		var columna = 0
		for potenciaDos < MODULO {
			var paridad = false
			for fila := 0; fila < MODULO; fila++ {
				if matriz4096[fila][columna] {
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

	if introducir_error == "true" {
		fmt.Print("Introducir error...")
		fmt.Print("\n")
		s1 := rand.NewSource(time.Now().UnixNano()) // Setea la semilla
		for modulo := 0; modulo < cantidadModulos; modulo++ {
			r1 := rand.New(s1)
			introducirError := r1.Intn(2) // Calcula un random n, donde 0 <= n < 2
			if introducirError == 1 {
				posicionError := r1.Intn(4095) + 1 // Calcula un random n, donde 1 <= n < 4096
				//str := fmt.Sprintf("Mod: %d, Pos: %d, Bit: %d\n", modulo, posicionError, arrayModules[modulo][posicionError])
				//fmt.Print(str)
				arrayModules[modulo][posicionError] = !arrayModules[modulo][posicionError]
			}
		}
	}

	var textoProtegido string
	for k := 0; k < len(arrayModules); k++ {
		if len(textoProtegido) == 0 {
			textoProtegido = helpers.TransformarArreglo4096BooleanosToString(arrayModules[k])
		} else {
			textoProtegido = textoProtegido + helpers.TransformarArreglo4096BooleanosToString(arrayModules[k])
		}
	}

	textoProtegerGeneradoBytes = helpers.TransformarArregloModulos4096BooleanosToArregloBytes(arrayModules)
	//fmt.Print(textoProtegerGeneradoBytes)
	//fmt.Print("\n")

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoProtegido, textoProtegerGeneradoBytes, elapsed
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

	if corregir_error == "true" {
		fileAsBytes = CorregirError256(fileAsBytes)
	}

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

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoDesprotegido, textoDesprotegerGeneradoBytes, elapsed
}

func CorregirError256(archivoBytes []byte) []byte {
	archivoBooleano := helpers.TransformarArregloBytesToArregloBool(archivoBytes)

	cantModulos := helpers.CalcularCantidadModulos(archivoBooleano, constants.TAM_BITS_TOTALES_MODULO_256)

	fmt.Printf("Cantidad de modulos calculada: %d \n", cantModulos)

	arregloModulos := helpers.CrearArregloDeModulos256(archivoBooleano, cantModulos)

	matriz256 := helpers.GenerarMatriz256()

	for indiceModulo := 0; indiceModulo < cantModulos; indiceModulo++ {
		modulo := arregloModulos[indiceModulo]
		if posicionConError := helpers.ChequearErrorModulo256(modulo, matriz256); posicionConError != 0 {
			fmt.Printf("Se encontro error en el modulo %d y en la posicion %d \n", indiceModulo, posicionConError)
			helpers.CorregirErrorModulo256(arregloModulos, indiceModulo, posicionConError)
		}
	}

	return helpers.TransformarArregloModulos256BooleanosToArregloBytes(arregloModulos)
}

func DesprotegerHamming1024(ctx context.Context, fileName string, corregir_error string) (string, string, []byte, float64) {
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

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 1013, 10
	const MODULO = CANTIDAD_BITS_INFO + CANTIDAD_BITS_CONTROL + 1
	const CANTIDAD_BYTES_MODULO = MODULO / 8

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
	var cantidadBitsInformacion = cantidadModulos * 1013
	fmt.Print(cantidadBitsInformacion)
	fmt.Print("\n")

	byteObtenido := ""
	arrayBitsInformacion := make([]bool, cantidadBitsInformacion)
	indexModulo1024 := 0
	indexInfo := 0
	controlBytesModulo := 0

	if corregir_error == "true" {
		fileAsBytes = CorregirError1024(fileAsBytes)
	}

	for _, n := range fileAsBytes {
		byteObtenido = fmt.Sprintf("%08b", n)
		for _, value := range byteObtenido {
			// Si es un bit de control lo tenemos que ignorar
			var potenciaDos int
			potenciaDos = (indexModulo1024 & (indexModulo1024 - 1))
			if potenciaDos == 0 || indexModulo1024 == 0 {
				/*
					fmt.Print("Saltamos la pos : ")
					fmt.Print(indexModulo1024)
					fmt.Print("\n")
				*/
				indexModulo1024 = indexModulo1024 + 1
				continue
			}

			if string(value) == "1" {
				arrayBitsInformacion[indexInfo] = true
			} else {
				arrayBitsInformacion[indexInfo] = false
			}
			indexInfo = indexInfo + 1
			indexModulo1024 = indexModulo1024 + 1
		}

		controlBytesModulo = controlBytesModulo + 1

		// Si ya se consumió un modulo entero hay que reiniciar las variables
		if controlBytesModulo == CANTIDAD_BYTES_MODULO {
			indexModulo1024 = 0
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

	elapsed := time.Since(start).Seconds()
	//fmt.Printf("Tiempo transcurrido TOTAL: %s\n", elapsed)
	return string(fileAsBytes), textoDesprotegido, textoDesprotegerGeneradoBytes, elapsed
}

func CorregirError1024(archivoBytes []byte) []byte {
	archivoBooleano := helpers.TransformarArregloBytesToArregloBool(archivoBytes)

	cantModulos := helpers.CalcularCantidadModulos(archivoBooleano, constants.TAM_BITS_TOTALES_MODULO_1024)

	fmt.Printf("Cantidad de modulos calculada: %d \n", cantModulos)

	arregloModulos := helpers.CrearArregloDeModulos1024(archivoBooleano, cantModulos)

	matriz1024 := helpers.GenerarMatriz1024()

	for indiceModulo := 0; indiceModulo < cantModulos; indiceModulo++ {
		modulo := arregloModulos[indiceModulo]
		if posicionConError := helpers.ChequearErrorModulo1024(modulo, matriz1024); posicionConError != 0 {
			fmt.Printf("Se encontro error en el modulo %d y en la posicion %d \n", indiceModulo, posicionConError)
			helpers.CorregirErrorModulo1024(arregloModulos, indiceModulo, posicionConError)
		}
	}

	return helpers.TransformarArregloModulos1024BooleanosToArregloBytes(arregloModulos)
}

func DesprotegerHamming2048(ctx context.Context, fileName string, corregir_error string) (string, string, []byte, float64) {
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

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 2036, 11
	const MODULO = CANTIDAD_BITS_INFO + CANTIDAD_BITS_CONTROL + 1
	const CANTIDAD_BYTES_MODULO = MODULO / 8

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
	var cantidadBitsInformacion = cantidadModulos * 2036
	fmt.Print(cantidadBitsInformacion)
	fmt.Print("\n")

	byteObtenido := ""
	arrayBitsInformacion := make([]bool, cantidadBitsInformacion)
	indexModulo2048 := 0
	indexInfo := 0
	controlBytesModulo := 0

	for _, n := range fileAsBytes {
		byteObtenido = fmt.Sprintf("%08b", n)
		for _, value := range byteObtenido {
			// Si es un bit de control lo tenemos que ignorar
			var potenciaDos int
			potenciaDos = (indexModulo2048 & (indexModulo2048 - 1))
			if potenciaDos == 0 || indexModulo2048 == 0 {
				/*
					fmt.Print("Saltamos la pos : ")
					fmt.Print(indexModulo2048)
					fmt.Print("\n")
				*/
				indexModulo2048 = indexModulo2048 + 1
				continue
			}

			if string(value) == "1" {
				arrayBitsInformacion[indexInfo] = true
			} else {
				arrayBitsInformacion[indexInfo] = false
			}
			indexInfo = indexInfo + 1
			indexModulo2048 = indexModulo2048 + 1
		}

		controlBytesModulo = controlBytesModulo + 1

		// Si ya se consumió un modulo entero hay que reiniciar las variables
		if controlBytesModulo == CANTIDAD_BYTES_MODULO {
			indexModulo2048 = 0
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

func CorregirError2048(archivoBytes []byte) []byte {
	archivoBooleano := helpers.TransformarArregloBytesToArregloBool(archivoBytes)

	cantModulos := helpers.CalcularCantidadModulos(archivoBooleano, constants.TAM_BITS_TOTALES_MODULO_1024)

	fmt.Printf("Cantidad de modulos calculada: %d \n", cantModulos)

	arregloModulos := helpers.CrearArregloDeModulos1024(archivoBooleano, cantModulos)

	matriz1024 := helpers.GenerarMatriz1024()

	for indiceModulo := 0; indiceModulo < cantModulos; indiceModulo++ {
		modulo := arregloModulos[indiceModulo]
		if posicionConError := helpers.ChequearErrorModulo1024(modulo, matriz1024); posicionConError != 0 {
			fmt.Printf("Se encontro error en el modulo %d y en la posicion %d \n", indiceModulo, posicionConError)
			helpers.CorregirErrorModulo1024(arregloModulos, indiceModulo, posicionConError)
		}
	}

	return helpers.TransformarArregloModulos1024BooleanosToArregloBytes(arregloModulos)
}

func DesprotegerHamming4096(ctx context.Context, fileName string, corregir_error string) (string, string, []byte, float64) {
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

	const CANTIDAD_BITS_INFO, CANTIDAD_BITS_CONTROL = 4083, 12
	const MODULO = CANTIDAD_BITS_INFO + CANTIDAD_BITS_CONTROL + 1
	const CANTIDAD_BYTES_MODULO = MODULO / 8

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
	var cantidadBitsInformacion = cantidadModulos * 4083
	fmt.Print(cantidadBitsInformacion)
	fmt.Print("\n")

	byteObtenido := ""
	arrayBitsInformacion := make([]bool, cantidadBitsInformacion)
	indexModulo4096 := 0
	indexInfo := 0
	controlBytesModulo := 0

	for _, n := range fileAsBytes {
		byteObtenido = fmt.Sprintf("%08b", n)
		for _, value := range byteObtenido {
			// Si es un bit de control lo tenemos que ignorar
			var potenciaDos int
			potenciaDos = (indexModulo4096 & (indexModulo4096 - 1))
			if potenciaDos == 0 || indexModulo4096 == 0 {
				/*
					fmt.Print("Saltamos la pos : ")
					fmt.Print(indexModulo4096)
					fmt.Print("\n")
				*/
				indexModulo4096 = indexModulo4096 + 1
				continue
			}

			if string(value) == "1" {
				arrayBitsInformacion[indexInfo] = true
			} else {
				arrayBitsInformacion[indexInfo] = false
			}
			indexInfo = indexInfo + 1
			indexModulo4096 = indexModulo4096 + 1
		}

		controlBytesModulo = controlBytesModulo + 1

		// Si ya se consumió un modulo entero hay que reiniciar las variables
		if controlBytesModulo == CANTIDAD_BYTES_MODULO {
			indexModulo4096 = 0
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
