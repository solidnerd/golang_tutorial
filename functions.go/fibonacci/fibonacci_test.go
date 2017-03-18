//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Samples for go syntax - calculate fibonacci number via recursion

package fibonacci

import (
	"testing"

	"github.com/framps/golang_tutorial/functions.go/fibonacci"
)

var fibonacciTests = []struct {
	input    int // function input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

// TestFibonacci -
func TestFibonacci(t *testing.T) {
	for _, tt := range fibonacciTests {
		t.Logf("Calculating Fibonacci number for %d", tt.input)
		actual, _ := fibonacci.Fibonacci(tt.input)
		if actual != tt.expected {
			t.Errorf("Fibonacci(%d): expected %d, actual %d", tt.input, tt.expected, actual)
		}
	}
}
