package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input    int
	expected int
}{
	{input: 2, expected: 2},
	{input: 10, expected: 5},
	{input: 13195, expected: 29},
}

func TestCalculateLargestPrimeNumber(t *testing.T) {
	for _, v := range table {
		t.Run(fmt.Sprintf("input_size: %d", v.input), func(t *testing.T) {
			sum := calculateLargestPrimeNumber(v.input)
			if v.expected != sum {
				t.Fail()
			}
		})
	}

}
