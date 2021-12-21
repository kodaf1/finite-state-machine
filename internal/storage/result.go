package storage

import (
	"github.com/kodaf1/finite-state-machine/internal/fms/states"
)

type Result map[int]Data

type Data map[states.State]states.State
