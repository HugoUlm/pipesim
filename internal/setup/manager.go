package setup

import (
    "fmt"
    "os/exec"
	"strings"

	"github.com/HugoUlm/pipesim/internal/types"
)

func InstallLanguage(setup *types.LanguageSetup, useCache bool) (string, error) {
	if isInstalled(setup.Language, setup.Version) {
		fmt.Printf("%s %s already installed, skipping...\n", strings.Title(setup.Language.String()), setup.Version)
		return "", nil
	}

    var cmd string

    switch setup.Language {
    case types.DotNet:
		cmd = fmt.Sprintf("brew install dotnet@%s", strings.Split(setup.Version, ".")[0])
	case types.GoLang:
        cmd = fmt.Sprintf(`brew install go@%s`, setup.Version)
	case types.NodeJS:
        fmt.Sprintf(`brew install nvm && \
			mkdir -p $HOME/.nvm && \
			export NVM_DIR="$HOME/.nvm" && \
			[ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && . "/opt/homebrew/opt/nvm/nvm.sh" || \
			[ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh" && \
			nvm install %s`, setup.Version)
	default:
        return "", fmt.Errorf("Unsupported language: %s", setup.Language.String())
    }

    return cmd, nil
}

func CleanupInstall(language string, useCache bool) error {
	if useCache {
		fmt.Println("✅ Won't clean up because of --use-cache")
		return nil
	}

	fmt.Println("🧹 Cleaning up...")
	exec.Command("bash", "-c", fmt.Sprintf("brew uninstall %s", language))
	return nil
}

func isInstalled(lang types.Language, version string) bool {
    switch lang {
    case types.DotNet:
        out, err := exec.Command("dotnet", "--list-sdks").Output()
        return err == nil && strings.Contains(string(out), version[:3])
    case types.GoLang:
        out, err := exec.Command("go", "version").Output()
        return err == nil && strings.Contains(string(out), version[:len("1.XX")])
    case types.NodeJS:
        out, err := exec.Command("node", "--version").Output()
        return err == nil && strings.Contains(string(out), version[:len("XX")])
    default:
        return false
    }
}
