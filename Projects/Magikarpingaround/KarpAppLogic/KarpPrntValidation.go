package CMD

import (
	"errors"
	"strconv"
	"strings"
)

func PrintValidation(id string) ([]Task, error) {
	var TemptaskTrack []Task
	if len(tasks) == 0 {
		return nil, errors.New("No saved Tasks currently")
	}
	TemptaskTrack = nil
	argtokens := strings.TrimSpace(id)
	if argtokens == "" {
		return tasks, nil
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
					return nil, errors.New("The input: " + stRang + " Is an invalid argument. Please Enter an ID or Type Check With No ID")
				}
			}

			for i := range tasks {
				if tasks[i].ID == usrIN {

					TemptaskTrack = append(TemptaskTrack, tasks[i])
					//"Task Number:", tasks[i].ID, "// Desc:", tasks[i].Description, "// Time added/Created:", tasks[i].Date, "// Priority:", tasks[i].Priority, "// Completion:", tasks[i].Completion)//
					boolflag = true

				}

			}
			if boolflag == false {
				if len(multiput) > 1 {
					continue
				} else {
					return nil, errors.New("The task with the ID of: " + strconv.Itoa(usrIN) + " could not be found.")
				}

			}
		}
	}
	return TemptaskTrack, nil

}
