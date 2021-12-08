package states

import (
	"errors"
)

type SymbolB struct {
}

func (s SymbolB) IsFinal() bool {
	return false
}

func (s SymbolB) Next(n rune) (State, error) {
	switch n {
	case 'b':
		return SymbolB{}, nil
	case 'c':
		return SymbolC{}, nil
	default:
		return nil, errors.New("incorrect syntax")
	}
}
