package pipeline

import (
	"os"
	"fmt"

	"gopkg.in/yaml.v3"
)

func Parse(filename string) (*Pipeline, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var workflow Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, err
	}

	var steps []Step
	for _, job := range workflow.Jobs {
		for _, s := range job.Steps {
			switch {
			case s.Run != "":
				steps = append(steps, Step{
					Name:    s.Name,
					Command: s.Run,
					Env:     s.Env,
				})
			case s.Uses == "actions/checkout@v4":
				steps = append(steps, Step{
					Name:    s.Name,
					Command: "echo 📦 Cloning repository (simulated)...",
				})
			case s.Uses != "":
				steps = append(steps, Step{
					Name:    s.Name,
					Command: fmt.Sprintf("echo ⚠️ Skipping unsupported action: %s", s.Uses),
				})
			default:
				continue
			}
		}
		break // simulate only the first job for now
	}

	return &Pipeline{Steps: steps}, nil
}
