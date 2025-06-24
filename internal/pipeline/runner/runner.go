package runner

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/HugoUlm/pipesim/internal/types"
	"github.com/HugoUlm/pipesim/internal/setup"
)

func Run(commands []types.Command, dryRun bool, useCache bool) {
	var language string
	for _, command := range commands {
		fmt.Printf("▶️ %s\n", command.Name)

		if dryRun {
			fmt.Printf("    would run: %s\n", command.Cmd)
			continue
		}

		language = command.Language
		var cmd *exec.Cmd

		if command.Project != "" {
			cmd = exec.Command("bash", "-c", fmt.Sprintf("%s %s", command.Cmd, command.Project))
		} else {
			cmd = exec.Command("bash", "-c", fmt.Sprintf("%s", command.Cmd))
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Set env
		if len(command.Env) > 0 {
			env := os.Environ()
			for k, v := range command.Env {
				env = append(env, fmt.Sprintf("%s=%s", k, v))
			}
			cmd.Env = env
		}

		start := time.Now()
		err := cmd.Run()
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("❌ failed (%v)\n", elapsed)
		} else {
			fmt.Printf("✅ succeeded (%v)\n", elapsed)
		}
	}

	setup.CleanupInstall(language, useCache)
}
