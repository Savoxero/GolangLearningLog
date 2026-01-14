package CMD

import (
	"fmt"
)

func Print(input []Task) {
	if len(input) > 1 {
		fmt.Println(input)
	} else if len(input) == 1 {
		fmt.Println(input)
	}
	// to do : format the printing messages of both add and delete
}
