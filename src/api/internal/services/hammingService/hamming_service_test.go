package hammingservice

import "testing"

func TestCalcularCantidadModulos(t *testing.T) {
	type args struct {
		fileAsBool  []bool
		cantModulos int
	}

	tests := []struct {
		name          string
		args          args
		expectedValue int
	}{
		{
			name: "Caso base",
			args: args{
				fileAsBool:  nil,
				cantModulos: 0,
			},
			expectedValue: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

/*
func TestCrearArregloBool(t *testing.T) {
	type args struct {
		arregloBytes []byte
	}

	tests := []struct {
		name          string
		args          args
		expectedValue []bool
	}{
		{
			name: "Test con 1 byte",
			args: args{
				arregloBytes: []byte{97},
			},
			expectedValue: convertirStringArregloBool("01100001"),
		},
		{
			name: "Test con 2 bytes",
			args: args{
				arregloBytes: []byte{97, 54},
			},
			expectedValue: convertirStringArregloBool("01100001" + "00110110"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CrearArregloBool(tt.args.arregloBytes)

			if !compararArreglosBool(result, tt.expectedValue) {
				t.Errorf("No son iguales - result: %v - expected: %v", result, tt.expectedValue)
			}
		})
	}
}



func TestCrearArregloDeModulos256(t *testing.T) {

}

func TestCalcularValorDecimal(t *testing.T) {
	type args struct {
		input []bool
	}

	tests := []struct {
		name          string
		args          args
		expectedValue int
	}{
		{
			name: "Cadena 0101",
			args: args{
				input: convertirStringArregloBool("0101"),
			},
			expectedValue: 5,
		},
		{
			name: "Cadena 0000",
			args: args{
				input: convertirStringArregloBool("0000"),
			},
			expectedValue: 0,
		},
		{
			name: "Cadena 1000",
			args: args{
				input: convertirStringArregloBool("1000"),
			},
			expectedValue: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalcularValorDecimal(tt.args.input)

			if result != tt.expectedValue {
				t.Errorf("Error. Expected value: %d -- Result value: %d", tt.expectedValue, result)
			}
		})
	}
}

func compararArreglosBool(arr1 []bool, arr2 []bool) bool {
	result := true
	if len(arr1) != len(arr2) {
		result = false
		return result
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			result = false
			break
		}
	}

	return result
}

func convertirStringArregloBool(str string) []bool {
	result := make([]bool, len(str))

	for i, n := range str {
		if string(n) == "1" {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

*/
