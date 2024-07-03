package stateManager

var States, Commands = InitStatesAndCommands()

type State struct {
	Name             string
	PossibleCommands []*Command
}

// well, this is an ugly function, but basically we "create" commands
// then we create a state that uses that commands
// then we create more commands, and one more state that uses that commands
// and then we atribute that state to the previous commands as they should
// (sorry for the paragraph)
func InitStatesAndCommands() ([]State, []Command) {
	existingStates := []State{}
	existingCommands := []Command{}

	//command 0
	existingCommands = append(existingCommands, Command{
		Name: "HELO",
	})
	//command 1
	existingCommands = append(existingCommands, Command{
		Name: "EHLO",
	})

	//state 0
	existingStates = append(existingStates, State{
		Name: "Initial",
		PossibleCommands: []*Command{
			&existingCommands[0],
			&existingCommands[1],
		},
	})

	//command 2
	existingCommands = append(existingCommands, Command{
		Name: "MAIL",
	})

	//state 1
	existingStates = append(existingStates, State{
		Name: "Comunication Started",
		PossibleCommands: []*Command{
			&existingCommands[2],
		},
	})

	//dar update nos commands 0 e 1
	existingCommands[0].NextState = &existingStates[1]
	existingCommands[1].NextState = &existingStates[1]

	//just for the sake of having continuation the last one will poit to the first until we implement quit command
	existingCommands[2].NextState = &existingStates[0]
	return existingStates, existingCommands
}
