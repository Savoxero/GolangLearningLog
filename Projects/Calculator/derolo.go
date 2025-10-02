package main

import "fmt"

func loli() { // Function at package level, NOT inside main
	const firstName = "Lane"
	const lastName = "Wagner"
	const fullName = firstName + " " + lastName
	fmt.Println(fullName)

}
func heybitches() {
	var loli int8 = 99
	fmt.Println(loli)
}
