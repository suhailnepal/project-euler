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
	{input: 10, expected: 10},
	{input: 100, expected: 44},
}

func TestFebonacciSum(t *testing.T) {
	for _, v := range table {
		t.Run(fmt.Sprintf("input_size: %d", v.input), func(t *testing.T) {
			sum := febonacciSum(v.input)
			if v.expected != sum {
				t.Fail()
			}
		})
	}

}
