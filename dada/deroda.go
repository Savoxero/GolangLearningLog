package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var operation string

	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(&operation)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	if operation == "-" {
		result := subtract(num1, num2)
		fmt.Println(result)
		return
	} else if operation == "*" {
		result := multiply(num1, num2)
		fmt.Println(result)
		return
	} else if operation == "/" {
		result, nil := divide(num1, num2)
		fmt.Println(result, nil)
		return
	} else if operation == "+" {
		result := add(num1, num2)
		fmt.Println(result)

		return

	}
}
