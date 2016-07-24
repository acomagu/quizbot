package main

import ()

type State struct {
	qa QA
	stage int
}

var states map[string]State = make(map[string]State)

func setStateBy(userID string, state State) {
	states[userID] = state
}

func stateBy(userID string) State {
	state, ok := states[userID]
	if !ok {
		state = initialState()
	}
	return state
}

func initialState() State {
	return State{
		qa: oneQA(),
		stage: 0,
	}
}
