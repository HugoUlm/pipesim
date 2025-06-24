package types

type Command struct {
	Name		string
	Cmd			string
	Env			map[string]string
	Language	string
	Project		string
}
