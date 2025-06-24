package types

type Pipeline struct {
	Steps []Step `yaml:"steps"`
}

type Workflow struct {
    Name string         `yaml:"name"`
    On   Trigger         `yaml:"on"`
    Jobs map[string]Job  `yaml:"jobs"`
}

type Trigger struct {
    Push            BranchFilter `yaml:"push,omitempty"`
    PullRequest     BranchFilter `yaml:"pull_request,omitempty"`
    WorkflowDispatch interface{} `yaml:"workflow_dispatch,omitempty"`
}

type BranchFilter struct {
    Branches []string `yaml:"branches,omitempty"`
}

type Job struct {
    Name     string             `yaml:"name,omitempty"`
    RunsOn   string             `yaml:"runs-on"`
    Needs    interface{}        `yaml:"needs,omitempty"`
    If       string             `yaml:"if,omitempty"`
    Strategy *Strategy          `yaml:"strategy,omitempty"`
    Steps    []Step             `yaml:"steps"`
    Env      map[string]string  `yaml:"env,omitempty"`
}

type Strategy struct {
    Matrix map[string][]string `yaml:"matrix,omitempty"`
}

type Step struct {
    Name string `yaml:"name,omitempty"`
    Uses string `yaml:"uses,omitempty"`
    Run  string `yaml:"run,omitempty"`
    With map[string]string `yaml:"with,omitempty"`
    Env  map[string]string `yaml:"env,omitempty"`
}

