package types

type Pipeline struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name            string            `yaml:"name"`
	Command         string            `yaml:"command"`
	Env             map[string]string `yaml:"env"`
	ContinueOnError bool              `yaml:"continueOnError"`
	Uses			string		  `yaml:"uses"`
	Run				string		  `yaml:"run"`
	With			map[string]string `yaml:"with"`
}

type Workflow struct {
    Name string         `yaml:"name"`
    On   Trigger        `yaml:"on"`
    Jobs map[string]Job `yaml:"jobs"`
}

type Trigger struct {
    Push        BranchFilter `yaml:"push,omitempty"`
    PullRequest BranchFilter `yaml:"pull_request,omitempty"`
}

type BranchFilter struct {
    Branches []string `yaml:"branches,omitempty"`
}

type Job struct {
    Strategy *Strategy `yaml:"strategy,omitempty"`
    Steps    []Step    `yaml:"steps,omitempty"`
}

type Strategy struct {
    Matrix map[string][]string `yaml:"matrix,omitempty"`
}
