package main

import (
	"fmt"
	"testing"
)

var table = []struct {
	input    int
	expected int
}{
	{input: 2, expected: 0},
	{input: 10, expected: 23},
	{input: 100, expected: 2318},
}

func TestFindTotalSumOfMultiples(t *testing.T) {
	for _, v := range table {
		t.Run(fmt.Sprintf("input_size: %d", v.input), func(t *testing.T) {
			sum := FindTotalSumOfMultiples(v.input)
			if v.expected != sum {
				t.Fail()
			}
		})
	}

}
