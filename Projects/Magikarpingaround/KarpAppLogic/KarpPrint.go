package CMD

import (
	"fmt"
	"strconv"
	"strings"
)

func Print(id string) {
	argtokens := strings.TrimSpace(id)
	multiput := strings.Split(argtokens, "&")
	if len(tasks) == 0 && argtokens == "" {
		fmt.Println("no saved tasks currently")

	}
	if argtokens != "" {
		for _, stRang := range multiput {
			boolflag := false
			stRang = strings.TrimSpace(stRang)
			usrInerr := stRang
			dada, err := strconv.Atoi(stRang)
			if err != nil {
				fmt.Println(usrInerr, "Is an invalid argument. Please Enter an ID or Type Check With No ID")
				continue
			}
			for i := range tasks {
				if tasks[i].ID == dada {
					fmt.Println("Task Number:", tasks[i].ID, "// Desc:", tasks[i].Description, "// Time added/Created:", tasks[i].Date, "// Priority:", tasks[i].Priority, "// Completion:", tasks[i].Completion)
					boolflag = true
					continue
				}

			}
			if boolflag == false {
				fmt.Println("The task with the ID of", dada, "could not be found.")
				continue
			}
		}
	}
	if argtokens == "" {
		fmt.Println(tasks)

	}
}
