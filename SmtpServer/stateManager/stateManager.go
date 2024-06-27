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

func (sm *StateManager) NextState() {
	sm.CurrentStateIndex++
	sm.CurrentState = States[sm.CurrentStateIndex]
}
