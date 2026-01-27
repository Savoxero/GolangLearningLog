package CMD

import (
	Domain "MagikarpingAround/KarpDom"
	"errors"
	"strconv"
	"strings"
)

func PrintValidation(id string) ([]Domain.Task, error) {
	var TemptaskTrack []Domain.Task
	argtokens := strings.TrimSpace(id)
	if argtokens == "" {
		var emptyinp []int
		TemptaskTrack, err := Domain.KarpPrint(emptyinp)
		if err != nil {
			return nil, err // in case tasks is empty
		} else {
			return TemptaskTrack, nil
		}
	}
	multiput := strings.Split(argtokens, "&")
	if argtokens != "" {
		for _, STRrange := range multiput {
			STRrange = strings.TrimSpace(STRrange)
			usrIN, err := strconv.Atoi(STRrange)
			if err != nil {
				return nil, errors.New("The input: " + STRrange + " Is an invalid argument. Please Enter an ID or Type Check With No ID")
			}
			var MultID []int
			MultID = append(MultID, usrIN)
			TemptaskTrack, err = Domain.KarpPrint(MultID)
			if err != nil {
				return nil, err
			}
		}

	}
	return TemptaskTrack, nil

}
