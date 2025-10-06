package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 float64
	var num2 float64
	var operation string
	var input string
	var filterednumb float64
	var err error
	//var finalnumresult float64///

	for {
		fmt.Print("Enter first number: ")
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("thank you, come again")
			return
		}

		filterednumb, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please Enter a valid number")
			continue
		}

		num1 = filterednumb

		for {
			fmt.Print("Enter operation (+, -, *, /): ")
			fmt.Scan(&operation)
			if operation == "exit" {
				fmt.Println("thank you, come again")
				return
			}
			if operation == "+" || operation == "/" || operation == "-" || operation == "*" || operation == "exit" {
				break
			}
			if operation != "+" && operation != "/" && operation != "-" && operation != "*" && operation != "exit" {
				fmt.Println("Please Enter a valid operation")
				continue
			}

			filterednumb, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Please Enter a valid operation")
				break
			}
		}

		for {
			fmt.Print("Enter second number: ")
			fmt.Scan(&input)
			if input == "exit" {
				return
			}
			filterednumb, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Please Enter a valid number")
				continue
			}
			break
		}

		num2 = filterednumb

		switch operation {
		case "-":
			fmt.Println(subtract(num1, num2))
		//	finalnumresult = subtract(num1, num2)
		case "*":
			fmt.Println(multiply(num1, num2))
		///	finalnumresult = multiply(num1, num2)
		case "/":
			result, err := divide(num1, num2)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(result)
		///	finalnumresult = result
		case "+":
			fmt.Println(add(num1, num2))
		///	finalnumresult = add(num1, num2)

		default:
			fmt.Println("Unknown operation")
		}
		///finalnumresult = num1

	}

}
