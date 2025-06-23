package pipeline

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Run(p *Pipeline, dryRun bool, project string) {
	for _, step := range p.Steps {
		fmt.Printf("▶ %s\n", step.Name)

		if dryRun {
			fmt.Printf("    would run: %s\n", step.Command)
			continue
		}

		cmd := exec.Command("bash", "-c", fmt.Sprintf("%s %s", step.Command, project))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Set env
		if len(step.Env) > 0 {
			env := os.Environ()
			for k, v := range step.Env {
				env = append(env, fmt.Sprintf("%s=%s", k, v))
			}
			cmd.Env = env
		}

		start := time.Now()
		err := cmd.Run()
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("✘ %s failed (%v)\n", step.Name, elapsed)
			if !step.ContinueOnError {
				break
			}
		} else {
			fmt.Printf("✔ %s succeeded (%v)\n", step.Name, elapsed)
		}
	}
}
