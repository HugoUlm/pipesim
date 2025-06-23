package pipeline

type Pipeline struct {
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name            string            `yaml:"name"`
	Command         string            `yaml:"command"`
	Env             map[string]string `yaml:"env"`
	ContinueOnError bool              `yaml:"continueOnError"`
	Uses		string		  `yaml:"uses"`
	Run		string		  `yaml:"run"`
}

type Workflow struct {
	Jobs map[string]struct {
		Steps []Step `yaml:"steps"`
	} `yaml:"jobs"`
}
