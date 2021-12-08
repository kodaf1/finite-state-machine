package states

import (
	"errors"
)

type SymbolC struct {
}

func (s SymbolC) IsFinal() bool {
	return false
}

func (s SymbolC) Next(n rune) (State, error) {
	switch n {
	case 'a':
		return SymbolA{}, nil
	case 'b':
		return SymbolB{}, nil
	case 'd':
		return SymbolD{}, nil
	default:
		return nil, errors.New("incorrect syntax")
	}
}
