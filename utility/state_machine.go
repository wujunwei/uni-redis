package utility

import "errors"

var(
	StateNotFound = errors.New("state not found in the machine")
)

type State interface {
	String() string
	IsEnd() bool
	IsDefault() bool
	From() State
	//Arg() interface{}
}

type Process func(interface{}) (State, error)

type StateMachine struct {
	currentState State
	lastState    State
	stateTable   map[string]Process
}



func (stateMachine *StateMachine) Register(from State, process Process) {
	if stateMachine.stateTable == nil {
		stateMachine.stateTable = make(map[string]Process)
	}
	stateMachine.stateTable[from.String()] = process
}

func (stateMachine *StateMachine) Current() State {
	return stateMachine.currentState
}

func (stateMachine *StateMachine) Last() State {
	return stateMachine.lastState
}

func (stateMachine *StateMachine) Fire(arg interface{}) (State, error) {
	 process , ok:= stateMachine.stateTable[stateMachine.currentState.String()]
	if ok {
		return process(arg)
	}
	return nil, StateNotFound
}
