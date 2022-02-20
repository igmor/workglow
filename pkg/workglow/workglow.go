package workglow

import (
	"os"

	"gopkg.in/yaml.v3"
)

type WorkGlow struct {
	Name  string
	Steps []interface{}
}

func ParseWorkGlow(workglowPath string) (*WorkGlow, error) {
	var rawYaml, err = os.ReadFile(workglowPath)
	if err != nil {
		return nil, err
	}

	var w WorkGlow
	if err = yaml.Unmarshal(rawYaml, &w); err != nil {
		return nil, err
	}
	return &w, nil
}
