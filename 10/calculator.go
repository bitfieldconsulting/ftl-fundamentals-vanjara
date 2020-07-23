package calculator

import "errors"

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Divide by zero")
	}
	return a / b, nil
}
