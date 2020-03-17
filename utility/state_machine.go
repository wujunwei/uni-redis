package utility

import "errors"

var (
	StateNotFound = errors.New("state not found in the machine")
)

type State interface {
	String() string
	IsEnd() bool
}

type DefaultState struct {
}

func (ds DefaultState) String() string {
	return "default"
}
func (ds DefaultState) IsEnd() bool {
	return false
}

type Process func(interface{}) (State, error)

// Deprecated: should not use it.
type StateMachine struct {
	currentState State
	lastState    State
	stateTable   map[string]Process
}

func (stateMachine *StateMachine) Register(state string, process Process) {
	if stateMachine.stateTable == nil {
		stateMachine.stateTable = make(map[string]Process)
	}
	stateMachine.stateTable[state] = process
}

func (stateMachine *StateMachine) Current() State {
	return stateMachine.currentState
}

func (stateMachine *StateMachine) Last() State {
	return stateMachine.lastState
}

func (stateMachine *StateMachine) Fire(arg interface{}) (State, error) {
	process, ok := stateMachine.stateTable[stateMachine.currentState.String()]
	if ok {
		return process(arg)
	}
	return nil, StateNotFound
}

func (stateMachine *StateMachine) Finished() bool {
	return stateMachine.currentState.IsEnd()
}

func newStateMachine(start State) *StateMachine {
	if nil == start {

	}
	return &StateMachine{currentState: start}
}
