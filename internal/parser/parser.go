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
        var variants []map[string]string

        if job.Strategy != nil && len(job.Strategy.Matrix) > 0 {
            variants = expandMatrix(job.Strategy.Matrix)
        } else {
            variants = []map[string]string{{}} // single default variant
        }

        for _, variant := range variants {
            for _, s := range job.Steps {
                // Apply matrix substitution
                s.Name = substituteMatrixVars(s.Name, variant)
                s.Run = substituteMatrixVars(s.Run, variant)
                s.Uses = substituteMatrixVars(s.Uses, variant)
                for k, v := range s.With {
                    s.With[k] = substituteMatrixVars(v, variant)
                }

                switch {
                case s.Run != "":
					projectValue := ""
					if stepRequiresProject(s) {
						projectValue = project
					}
                    commands = append(commands, types.Command{
                        Name:    s.Name,
                        Cmd:     s.Run,
                        Env:     s.Env,
                        Project: projectValue,
                    })

                case strings.Contains(s.Uses, "actions/setup-"):
                    version := resolveGoVersion(s, variant)
                    language := &types.LanguageSetup{
                        Language: types.ParseLanguage(strings.TrimPrefix(strings.Split(s.Uses, "@")[0], "actions/setup-")),
                        Version:  version,
                    }
                    cmd, _ := setup.InstallLanguage(language, useCache)

                    commands = append(commands, types.Command{
                        Name:     s.Name,
                        Cmd:      cmd,
                        Language: language.Language.String(),
                    })

                case strings.HasPrefix(s.Uses, "actions"):
                    commands = append(commands, types.Command{
                        Name: s.Name,
                        Cmd:  fmt.Sprintf(`echo "📦 simulating running %s"`, s.Uses),
                    })

                case s.Uses != "":
                    commands = append(commands, types.Command{
                        Name: s.Name,
                        Cmd:  fmt.Sprintf(`echo "⚠️ Skipping unsupported action: %s"`, s.Uses),
                    })

                default:
                    continue
                }
            }
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

func stepRequiresProject(step types.Step) bool {
    run := strings.ToLower(step.Run)
    return strings.Contains(run, "build") || strings.Contains(run, "test")
}
