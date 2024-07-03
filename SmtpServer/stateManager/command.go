package stateManager

type Command struct {
	Name      string
	NextState *State
}
