package utils

import (
	"fmt"
	"testing"
)

func Test_CalculateParityBitsNeeded(t *testing.T){
	tests := []struct{
		input string
		expected int
	}{
		{
			"1001",
			3,
		},
		{
			"10011",
			4,
		},
	}

	for _,tt := range tests {
		t.Run(fmt.Sprintf("Test CalculateParityBitsNeeded with input %s",tt.input), func(t *testing.T) {
			inputSize := len(tt.input)
			result := CalculateParityBitsNeeded(inputSize)

			if result != tt.expected {
				t.Errorf("Test_CalculateParityBitsNeeded result = %d, expected %d", result, tt.expected)
			}
		})
	}
}