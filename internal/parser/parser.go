package parser

import (
	"os"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
	"github.com/HugoUlm/pipesim/internal/setup"
	"github.com/HugoUlm/pipesim/internal/types"
)

func Parse(filename string, useCache bool, project string) ([]types.Command, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var workflow types.Workflow
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, err
	}

	var commands []types.Command
	for _, job := range workflow.Jobs {
		matrixVariants := []map[string]string{{}}
		if job.Strategy != nil && len(job.Strategy.Matrix) > 0 {
		    matrixVariants = ExpandMatrix(job.Strategy.Matrix)
		}

		for _, variant := range matrixVariants {
		    for _, s := range job.Steps {
		        name := SubstituteMatrixVars(s.Name, variant)
		        cmd := SubstituteMatrixVars(s.Run, variant)
		        uses := SubstituteMatrixVars(s.Uses, variant)

		        switch {
		        case cmd != "":
		            commands = append(commands, types.Command{
		                Name:    name,
		                Cmd:     cmd,
		                Env:     s.Env,
		                Project: project,
		            })
		        case strings.Contains(uses, "actions/setup-"):
					step.Name = substituteMatrixVars(step.Name, variant)
					step.Uses = substituteMatrixVars(step.Uses, variant)
					for k, v := range step.With {
					    step.With[k] = substituteMatrixVars(v, variant)
					}
		            version := detectLanguageVersions(step)
		            installCmd, _ := setup.InstallLanguage(version, useCache)
		            commands = append(commands, types.Command{
		                Name:     name,
		                Cmd:      installCmd,
		                Language: version.Language.String(),
		            })
		        case strings.HasPrefix(uses, "actions"):
		            commands = append(commands, types.Command{
		                Name: name,
		                Cmd:  fmt.Sprintf(`echo "📦 simulating running %s"`, uses),
		            })
		        case uses != "":
		            commands = append(commands, types.Command{
		                Name: name,
		                Cmd:  fmt.Sprintf(`echo "⚠️ Skipping unsupported action: %s"`, uses),
		            })
		        default:
		            continue
		        }
		    }
		    break // simulate only the first job for now
		}
	}
	return commands, nil
}

func detectLanguageVersions(step types.Step) *types.LanguageSetup {
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
