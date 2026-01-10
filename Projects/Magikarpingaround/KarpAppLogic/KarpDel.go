package CMD

import (
	"errors"
	"strconv"
	"strings"
)

func Del(argtokens string) ([]Task, error) {
	var lastdeltask []Task
	if len(tasks) == 0 {
		return nil, errors.New("No tasks currently saved.")
	}
	if len(argtokens) == 0 {
		return nil, errors.New("Empty input detected, Please enter a valid ID")

	}

	multdel := strings.Split(argtokens, "&")
	for _, IDIN := range multdel {
		found := false
		IDIN = strings.TrimSpace(IDIN)
		if IDIN == "" {
			continue
		}

		id, err := strconv.Atoi(IDIN)
		if err != nil {
			if len(lastdeltask) > 0 {
				continue
			} else {
				return nil, errors.New("Couldn't verify the input, Please enter a valid ID")
			}
		}

		for i, d := range tasks {
			if id == d.ID {
				lastdeltask = append(lastdeltask, tasks[i])
				found = true
				tasks = append(tasks[:i], tasks[i+1:]...)
				break

			}
		}

		if !found {
			if len(lastdeltask) > 0 {
				continue
			} else {
				return nil, errors.New("The number with the ID of " + IDIN + " Has not been found")
			}
		}
	}
	return lastdeltask, nil

}
