package states

import (
	"errors"
)

type Start struct {
}

func (s Start) IsFinal() bool {
	return false
}

func (s Start) Next(n rune) (State, error) {
	switch n {
	case 'a':
		return SymbolA{}, nil
	case 'b':
		return SymbolB{}, nil
	case 'c':
		return SymbolC{}, nil
	case 'd':
		return SymbolD{}, nil
	default:
		return nil, errors.New("incorrect syntax")
	}
}
