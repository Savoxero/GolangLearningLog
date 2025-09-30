package main

import "fmt"

func main() {
	result := add(16.5555, 19.2)
	fmt.Println(result)

	result2 := subtract(19, 98)
	fmt.Println(result2)

	result3, err := divide(24, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result3)
	}

	result4 := multiply(22, 44)
	fmt.Println(result4)
}
