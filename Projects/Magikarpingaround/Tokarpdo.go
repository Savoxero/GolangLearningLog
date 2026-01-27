package main

import (
	CMD "MagikarpingAround/KarpAppLogic"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// data, _, jaja := CMD.LoadFromfile()
	// fmt.Println(jaja)
	// fmt.Println(data)
	// CMD.WelcomeMSG()

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
			if Validation != nil { // double print, it's good enough semantics, if the user adds some empty input it gets returned as an error, and if the task was valid to be saved
				// it prints it. there could be a better alternative cause this looks a little redundant.
				fmt.Println(Validation)
				if len(Taskcreated) > 0 {
					fmt.Println("[Succesfully created the following Task/s:]")
					CMD.Print(Taskcreated)
				}

			} else if len(Taskcreated) > 0 {
				fmt.Println("[Succesfully created the following Task/s:]")
				CMD.Print(Taskcreated)
			}

		case "delete":
			TaskdeletionMsg, errwarning := CMD.Del(argtokens)
			if errwarning != nil {
				CMD.Print(TaskdeletionMsg)
				fmt.Println(errwarning)
				break
			}
			fmt.Println("[deleted the following Task/s:]")
			CMD.Print(TaskdeletionMsg)

		case "modify":
		case "print":
			ValidatedInput, PrntError := CMD.PrintValidation(argtokens)
			if PrntError != nil {
				fmt.Println(PrntError)
				break
			}
			CMD.Print(ValidatedInput)

		case "default":
		case "exit":
			fmt.Println("Goodbye, have a good day/night")
			return
		}
	}
}
