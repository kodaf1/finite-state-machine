package fms

import "github.com/kodaf1/finite-state-machine/internal/fms/states"

func Run(cmd string) bool {
	var state states.State
	state = states.Start{}

	for _, c := range cmd {
		newState, err := state.Next(c)
		if err != nil {
			return false
		}
		state = newState
	}

	return true
}
