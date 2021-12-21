package storage

import (
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
)

type Data map[states.State]states.State

func (d Data) Present() map[string]string {
	presentationMap := make(map[string]string)
	for k, v := range d {
		presentationMap[states.GetStateName(k)] = states.GetStateName(v)
	}
	return presentationMap
}
