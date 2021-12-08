package states

import (
	"errors"
)

type SymbolA struct {
}

func (s SymbolA) IsFinal() bool {
	return false
}

func (s SymbolA) Next(n rune) (State, error) {
	switch n {
	case 'c':
		return SymbolC{}, nil
	default:
		return nil, errors.New("incorrect syntax")
	}
}
