package parser

import (
	"fmt"
	"strings"

	"github.com/HugoUlm/pipesim/internal/types"
)

func expandMatrix(matrix map[string][]string) []map[string]string {
    if len(matrix) == 0 {
        return []map[string]string{}
    }

    // Generate Cartesian product of all matrix keys
    keys := make([]string, 0, len(matrix))
    for k := range matrix {
        keys = append(keys, k)
    }

    var recurse func(int, map[string]string) []map[string]string
    recurse = func(i int, partial map[string]string) []map[string]string {
        if i == len(keys) {
            // Copy result
            combination := make(map[string]string)
            for k, v := range partial {
                combination[k] = v
            }
            return []map[string]string{combination}
        }

        key := keys[i]
        var result []map[string]string
        for _, val := range matrix[key] {
            partial[key] = val
            result = append(result, recurse(i+1, partial)...)
        }
        return result
    }

    return recurse(0, make(map[string]string))
}

func resolveGoVersion(step types.Step, variant map[string]string) string {
	raw := step.With["go-version"]
	switch {
	case raw == "":
		return variant["go-version"]
	case strings.Contains(raw, "${{"):
		return substituteMatrixVars(raw, variant)
	default:
		return raw
	}
}

func substituteMatrixVars(input string, matrix map[string]string) string {
    for k, v := range matrix {
        placeholder := fmt.Sprintf("${{ matrix.%s }}", k)
        input = strings.ReplaceAll(input, placeholder, v)
    }
    return input
}
