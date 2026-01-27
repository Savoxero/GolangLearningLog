package CMD

import (
	Domain "MagikarpingAround/KarpDom"
	"strings"
)

func Add(argtokens string) ([]Domain.Task, error) {
	argtokens = strings.TrimSpace(argtokens)
	descriptions := strings.Split(argtokens, "&")
	for i, desc := range descriptions {
		descriptions[i] = strings.TrimSpace(desc)
		if desc == "" {
			continue
		}
	}
	taskcreationMSG, err := Domain.AddTask(descriptions)
	if err != nil {
		return taskcreationMSG, err
	}
	return taskcreationMSG, nil

}
