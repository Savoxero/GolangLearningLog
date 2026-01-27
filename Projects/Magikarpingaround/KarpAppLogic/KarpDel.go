package CMD

import (
	Domain "MagikarpingAround/KarpDom"
	"errors"
	"strconv"
	"strings"
)

func Del(argtokens string) ([]Domain.Task, error) {
	var ValidatedDEL []Domain.Task
	var IDholder []int
	multdel := strings.Split(argtokens, "&")
	for _, IDIN := range multdel {
		IDIN = strings.TrimSpace(IDIN)
		id, err := strconv.Atoi(IDIN)
		if err != nil {
			if len(multdel) > 2 {
				continue
			} else {
				return nil, errors.New("Couldn't verify the input, Please enter a valid ID")
			}
		}
		IDholder = append(IDholder, id)
		Validation, err := Domain.DeleteTask(IDholder)
		ValidatedDEL = append(ValidatedDEL, Validation...)
		if err != nil {
			return ValidatedDEL, err
		}

	}
	return ValidatedDEL, nil
}
