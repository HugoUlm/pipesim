package types

type LanguageSetup struct {
    Language Language
    Version  string
    Path     string
}

type Language int

const (
    UnknownLang Language = iota
    DotNet
    GoLang
    NodeJS
)

func (l Language) String() string {
    switch l {
    case DotNet:
        return "dotnet"
    case GoLang:
        return "go"
    case NodeJS:
        return "node"
    default:
        return "unknown"
    }
}

func ParseLanguage(s string) Language {
    switch s {
    case "dotnet":
        return DotNet
    case "go":
        return GoLang
    case "node":
        return NodeJS
    default:
        return UnknownLang
    }
}
