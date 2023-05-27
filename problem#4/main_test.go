package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input    int
	expected int
}{
	{input: 101110, expected: 101101},
	{input: 800000, expected: 793397},
}

func TestCalculateLargestPalindrome(t *testing.T) {
	for _, v := range table {
		t.Run(fmt.Sprintf("input_size: %d", v.input), func(t *testing.T) {
			sum, _ := CalculateLargestPalindrome(v.input)
			if v.expected != sum {
				t.Fail()
			}
		})
	}
}
