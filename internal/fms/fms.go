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
