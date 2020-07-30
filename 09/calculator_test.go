package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 0},
		{a: 3, b: 2, want: 1},
		{a: 6, b: 7, want: -1},
		{a: 6, b: -7, want: 13},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Subtract(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 3, b: 2, want: 6},
		{a: 6, b: 7, want: 42},
	}
	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Multiply(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 1},
		{a: 4, b: 2, want: 2},
		{a: 42, b: 7, want: 6},
	}
	for _, testCase := range testCases {
		got := calculator.Divide(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
