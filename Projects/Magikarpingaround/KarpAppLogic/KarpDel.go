package CMD

import (
	"errors"
	"strconv"
	"strings"
)

func Del(argtokens string) ([]string, error) {
	multdel := strings.Split(argtokens, "&")
	for _, IDIN := range multdel {
		found := false
		IDIN = strings.TrimSpace(IDIN)
		if IDIN == "" {
			return multdel, errors.New("Empty input detected. Enter a valid ID")
		}
		if len(tasks) == 0 {
			return multdel, errors.New("No tasks currently saved.")
		}
		id, err := strconv.Atoi(IDIN)
		if err != nil {
			return multdel, errors.New("Couldn't verify the ID, Please enter a valid ID")
		}

		for i, d := range tasks {
			if id == d.ID {
				found = true
				tasks = append(tasks[:i], tasks[i+1:]...)
				continue
			}

		}
		if found == false {
			id2 := string(id)
			return multdel, errors.New("the task with the ID of", id2, " could not be found")
		}

	}

}
