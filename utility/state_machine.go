package utility

type State interface {
	String() string
	IsEnd() bool
	IsDefault() bool
	From() State
	//Arg() interface{}
}

type Event string

type StateMachine struct {
	currentState State
	lastState    State
}

func (stateMachine *StateMachine) addEvent(from State, event Event, to State) {

}
