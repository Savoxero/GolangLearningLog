package main

import "errors"

func add(a, b float64) float64 {
	c := a + b
	return c
}

func subtract(a, b float64) float64 {
	c := a - b
	return c
}

func multiply(a, b float64) float64 {
	c := a * b
	return c
}
func divide(a, b float64) (float64, error) {

	if b == 0 {

		return 0, errors.New("cannot divide by zero")

	}
	return a / b, nil

}
