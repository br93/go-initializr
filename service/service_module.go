package service

import (
	"fmt"

	"go-initializr/common"
)

type Service struct{}

func (*Service) GetModules(args ...string) ([]string, error) {

	if len(args) == 0 {
		return nil, fmt.Errorf("argument invalid")
	}

	var modules []string

	for _, element := range args {
		result := common.Find(element)
		if string(result) == "" {
			return nil, fmt.Errorf("module not found")
		}
		modules = append(modules, fmt.Sprintf("%s %s", "go get", common.Output(result)))
	}

	return modules, nil
}
