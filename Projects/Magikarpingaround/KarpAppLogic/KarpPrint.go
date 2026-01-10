package CMD

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Print(id string) error {
	if len(tasks) == 0 {
		return errors.New("No saved Tasks currently")

	}
	argtokens := strings.TrimSpace(id)
	if argtokens == "" {
		fmt.Println(tasks)
	}
	multiput := strings.Split(argtokens, "&")
	if argtokens != "" {
		for _, stRang := range multiput {
			boolflag := false

			stRang = strings.TrimSpace(stRang)
			usrIN, err := strconv.Atoi(stRang)
			if err != nil {
				if len(multiput) > 1 {
					continue
				} else {
					return errors.New("The input: " + stRang + " Is an invalid argument. Please Enter an ID or Type Check With No ID")
				}
			}
			for i := range tasks {
				if tasks[i].ID == usrIN {
					fmt.Println("Task Number:", tasks[i].ID, "// Desc:", tasks[i].Description, "// Time added/Created:", tasks[i].Date, "// Priority:", tasks[i].Priority, "// Completion:", tasks[i].Completion)
					boolflag = true
					continue
				}

			}
			if boolflag == false {
				if len(multiput) > 1 {
					continue
				} else {
					return errors.New("The task with the ID of: " + strconv.Itoa(usrIN) + " could not be found.")
				}

			}
		}
	}
	return nil

}
