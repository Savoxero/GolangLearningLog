package main

import (
	logic "InfraStru/KarpPackages"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	data, _, jaja := logic.LoadFromfile()
	fmt.Println(jaja)
	fmt.Println(data)

	for {
		fmt.Print("> ")
		Valid := scanner.Scan()
		if !Valid {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "input ERROR:", err)
			} else {
				fmt.Println("input closed: (EOF)")
			}
			break
		}
		input := scanner.Text()
		filter := strings.Fields(input)
		if len(filter) <= 0 {
			fmt.Println("no input detected, enter help for a list of commands")
			continue
		}
		cmtokens := filter[0]
		argtokens := strings.Join(filter[1:], " ")
		input = cmtokens
		switch input {
		case "help":
			fmt.Println("")

		case "add":
			Taskcreated := logic.Add(argtokens)
			fmt.Println("Sucessfully created:", Taskcreated)
		case "remove":
			logic.Del(argtokens)

		case "modify":
		case "print":
			logic.Print(argtokens)

		case "default":
		case "exit":

			fmt.Println("Goodbye, have a good day/night")
			return
		}
	}
}
