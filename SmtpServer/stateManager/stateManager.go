package stateManager

type StateManager struct {
	Domain            string
	CurrentState      State
	CurrentStateIndex int
}

func InitStateManager() *StateManager {
	return &StateManager{
		Domain:            "",
		CurrentState:      States[0],
		CurrentStateIndex: 0,
	}
}

func (sm *StateManager) NextState(commandName string) {
	nextState := findNextStateFromCommandName(commandName)
	sm.CurrentState = *nextState
}

func (sm *StateManager) IsPossibleCommand(commandName string) bool {
	for _, command := range sm.CurrentState.PossibleCommands {
		if command.Name == commandName {
			return true
		}
	}

	return false
}

func findNextStateFromCommandName(commandName string) *State {
	for _, command := range Commands {
		if command.Name == commandName {
			return command.NextState
		}
	}

	return nil
}
