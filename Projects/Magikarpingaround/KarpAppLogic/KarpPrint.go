package CMD

import (
	Domain "MagikarpingAround/KarpDom"
	"fmt"
)

func Print(input []Domain.Task) {
	if len(input) > 1 {
		fmt.Printf("+------+---------------------+-----------+------------+--------------------+\n| ID:  | Description:        | Priority: | Completed: | Date added:        |\n+------+---------------------+-----------+------------+--------------------+\n")
		for i := range input {
			fmt.Printf("| %3d  | %19s | %9s | %10t |%2v |\n%s", input[i].ID, input[i].Description, input[i].Priority, input[i].Completion, input[i].Date, "+------+---------------------+-----------+------------+--------------------+\n")

		}

	} else if len(input) == 1 {
		Tempholder := input
		Tempholder = append(Tempholder, input[0])
		fmt.Printf("+------+---------------------+-----------+------------+--------------------+\n| ID:  | Description:        | Priority: | Completed: | Date added:        |\n+------+---------------------+-----------+------------+--------------------+")
		fmt.Printf("\n| %3d  | %19s | %9s | %10t |%2v |\n%s", Tempholder[0].ID, Tempholder[0].Description, Tempholder[0].Priority, Tempholder[0].Completion, Tempholder[0].Date, "+------+---------------------+-----------+------------+--------------------+\n")
	}
	// to do : format the printing messages of both add and delete
}
