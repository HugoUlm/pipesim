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
	Language		string
}

type Workflow struct {
	Name string `yaml:"name"`
	On map[string]struct {
		Push map[string]struct {
			Branches []string `yaml:"branches"`
		} `yaml:"push"`
		PullRequest map[string]struct {
			Branches []string `yaml:"branches"`
		} `yaml:"pull_request"`
	} `yaml:"on"`
	Jobs map[string]struct {
		Steps []Step `yaml:"steps"`
	} `yaml:"jobs"`
}
