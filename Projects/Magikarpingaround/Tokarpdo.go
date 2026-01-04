package main

import (
	CMD "CMD/KarpAppLogic"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data, _, jaja := CMD.LoadFromfile()
	fmt.Println(jaja)
	fmt.Println(data)
	CMD.WelcomeMSG()

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

			Taskcreated, Validation := CMD.Add(argtokens)
			if Validation != nil {
				fmt.Println(Validation)
				break
			}
			fmt.Println("Sucessfully created:", Taskcreated)

		case "remove":
			CMD.Del(argtokens)

		case "modify":
		case "print":
			CMD.Print(argtokens)

		case "default":
		case "exit":

			fmt.Println("Goodbye, have a good day/night")
			return
		}
	}
}
