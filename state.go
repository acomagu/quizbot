package main

import ()

type Stage int

type State struct {
	qa QA
	stage Stage
}

const (
	_Ask Stage = iota
	_Answer
)

var states map[UserID]State = make(map[UserID]State)

func setStateBy(userID UserID, state State) {
	states[userID] = state
}

func stateBy(userID UserID) State {
	state, ok := states[userID]
	if !ok {
		state = initialState()
	}
	return state
}

func initialState() State {
	return State{
		qa: oneQA(),
		stage: _Ask,
	}
}
