package fms

import (
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
	"github.com/kodaf1/finite-state-machine/internal/storage"
)

func Run(cmd string) (states.State, error) {
	var state states.State
	state = states.Start{}

	for _, c := range cmd {
		newState, err := state.Next(c)
		if err != nil {
			return state, err
		}
		state = newState
	}

	return state, nil
}

func PartialRun(cmd string, isStart bool) storage.Data {
	result := storage.Data{}

	if isStart {
		result[states.GetStartState()] = states.GetStartState()
	} else {
		for _, v := range states.GetAllStartStates() {
			result[v] = v
		}
	}

	for _, symbol := range cmd {
		for key, actualState := range result {
			newState, err := actualState.Next(symbol)
			if err != nil {
				delete(result, key)
				continue
			}

			result[key] = newState
		}
	}

	return result
}
