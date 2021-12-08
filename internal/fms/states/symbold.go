package states

import (
	"errors"
)

type SymbolD struct {
}

func (s SymbolD) IsFinal() bool {
	return false
}

func (s SymbolD) Next(n rune) (State, error) {
	switch n {
	case 'e':
		return SymbolE{}, nil
	default:
		return nil, errors.New("incorrect syntax")
	}
}
