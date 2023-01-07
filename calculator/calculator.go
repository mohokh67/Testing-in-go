package calculator

import (
	"errors"
	"fmt"
)

type OperationType string

const (
	SumOperation      = OperationType("sum")
	SubtractOperation = OperationType("subtract")
	MultiplyOperation = OperationType("multiply")
)

func Sum(num1, num2 float64) float64 {
	return num1 + num2
}

func Subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func Multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

func Operate(operation string) func(num1, num2 float64) float64 {
	switch operation {
	case string(SumOperation):
		return Sum
	case string(SubtractOperation):
		return Subtract
	case string(MultiplyOperation):
		return Multiply
	default:
		return func(num1, num2 float64) float64 {
			fmt.Println("activity not supported")
			return 0
		}
	}
}
