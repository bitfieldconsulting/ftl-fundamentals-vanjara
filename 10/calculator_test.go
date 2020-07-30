package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 0},
		{a: 3, b: 2, want: 1},
		{a: 6, b: 7, want: -1},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
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
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 13},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a, b          int
		want          int
		errorExpected bool
	}{
		{a: 1, b: 1, want: 1, errorExpected: false},
		{a: 42, b: 7, want: 6, errorExpected: false},
		{a: 8, b: 4, want: 2, errorExpected: false},
		{a: 4, b: 0, want: 0, errorExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if tc.errorExpected != (err != nil) {
			t.Fatalf("Divide(%d, %d): unexpected error status", tc.a, tc.b)
		}
		if tc.want != got {
			t.Errorf("Divide(%d, %d): want %d, got %d", tc.a, tc.b, tc.want, got)
		}
	}
}
