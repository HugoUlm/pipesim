package pipeline

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(filename string) (*Pipeline, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var p Pipeline
	err = yaml.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
