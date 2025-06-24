package parser

import (
	"os"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
	pipetypes "github.com/HugoUlm/pipesim/internal/pipeline/types"
	"github.com/HugoUlm/pipesim/internal/setup"
	"github.com/HugoUlm/pipesim/internal/types"
)

func Parse(filename string, useCache bool, project string) ([]types.Command, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var workflow pipetypes.Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, err
	}

	var commands []types.Command
	for _, job := range workflow.Jobs {
		for _, s := range job.Steps {
			switch {
			case s.Run != "":
				commands = append(commands, types.Command{
					Name: s.Name,
					Cmd: s.Run,
					Env: s.Env,
					Project: project,
				})
			case strings.Contains(s.Uses, "actions/setup-"):
				version := detectLanguageVersions(s)
				cmd, _ := setup.InstallLanguage(version, useCache)
				commands = append(commands, types.Command{
					Name:	s.Name,
					Cmd: cmd,
					Language: version.Language.String(),
				})
			case strings.HasPrefix(s.Uses, "actions"):
				commands = append(commands, types.Command{
					Name:    s.Name,
					Cmd: fmt.Sprintf(`echo "📦 simulating running %s"`, s.Uses),
				})
			case s.Uses != "":
				commands = append(commands, types.Command{
					Name:    s.Name,
					Cmd: fmt.Sprintf(`echo "⚠️ Skipping unsupported action: %s"`, s.Uses),
				})
			default:
				continue
			}
		}
		break // simulate only the first job for now
	}

	return commands, nil
}

func detectLanguageVersions(step pipetypes.Step) *types.LanguageSetup {
    lang := strings.TrimPrefix(strings.Split(step.Uses, "@")[0], "actions/setup-")

    for key, val := range step.With {
        if strings.Contains(key, "version") {
            return &types.LanguageSetup{
                Language: types.ParseLanguage(lang),
                Version:  val,
            }
        }
    }

    return nil
}
