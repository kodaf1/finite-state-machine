package fms

import (
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
	"github.com/kodaf1/finite-state-machine/internal/storage"
)

func Run(cmd string) states.State {
	state := states.GetStartState()

	for _, c := range cmd {
		state = state.Next(c)
	}

	return state
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
			result[key] = actualState.Next(symbol)
			if result[key] == states.GetErrorState() {
				delete(result, key)
			}
		}
	}

	return result
}
