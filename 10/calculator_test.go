package calculator_test

import (
	"calculator"
	"fmt"
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
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
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
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
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
	testCases := []struct {
		a, b          int
		want          int
		errorExpected bool
	}{
		{a: 1, b: 1, want: 1, errorExpected: false},
		{a: 42, b: 7, want: 6, errorExpected: false},
		{a: 8, b: 4, want: 3, errorExpected: true},
		{a: 4, b: 0, want: 0, errorExpected: true},
	}
	for _, testCase := range testCases {
		got, err := calculator.Divide(testCase.a, testCase.b)
		if err == nil {
			if testCase.want != got && testCase.errorExpected != true {
				t.Errorf("Add(%d, %d): want %d, got %d, err %v", testCase.a, testCase.b, testCase.want, got, err)
			}
		} else {
			fmt.Printf("expected error; got %v\n", err)
		}
	}
}
