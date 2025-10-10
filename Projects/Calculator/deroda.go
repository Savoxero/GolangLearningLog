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
	var finalnumresult float64
	var resultA float64
	var calculationcheck = true

	for {

		for {
			if num1 > 0 {
				break
			}
			fmt.Print("Enter first number: ")
			fmt.Scan(&input)
			if input == "exit" {
				fmt.Println("thank you, come again")
				return
			}

			filterednumb, err = strconv.ParseFloat(input, 64)
			calculationcheck = false
			if err != nil {
				fmt.Println("Please Enter a valid number")
				continue
			}
			break
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

		if calculationcheck == true {
			switch operation {
			case "-":
				fmt.Println(subtract(finalnumresult, num2))

			case "*":
				fmt.Println(multiply(finalnumresult, num2))

			case "/":
				result, err := divide(finalnumresult, num2)
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Println(result)

			case "+":
				fmt.Println(add(finalnumresult, num2))

			default:
				fmt.Println("Unknown operation")
			}
		} else if calculationcheck == false {

			switch operation {
			case "-":
				fmt.Println(subtract(num1, num2))
				resultA = subtract(num1, num2)
			case "*":
				fmt.Println(multiply(num1, num2))
				resultA = multiply(num1, num2)
			case "/":
				result, err := divide(num1, num2)
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Println(result)
				resultA = result
			case "+":
				fmt.Println(add(num1, num2))
				resultA = add(num1, num2)

			default:
				fmt.Println("Unknown operation")
			}

			finalnumresult = resultA

			fmt.Printf("The starting number is now %f", finalnumresult)
		}

	}
}
