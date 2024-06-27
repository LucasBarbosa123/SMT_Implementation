package stateManager

var States = []State{
	{
		Name:             "Initial",
		PossibleCommands: []string{"HELO", "EHLO"},
	},
	{
		Name:             "Comunication Started",
		PossibleCommands: []string{"HELO", "EHLO"},
	},
}

type State struct {
	Name             string
	PossibleCommands []string
}
