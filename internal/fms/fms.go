package fms

import "github.com/kodaf1/finite-state-machine/internal/fms/states"

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

func PartialRun(cmd string) (map[states.State]states.State, error) {
	result := make(map[states.State]states.State, 0)

	for _, v := range states.GetAllStartStates() {
		result[v] = v
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

	return result, nil
}
