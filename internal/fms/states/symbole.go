package states

import (
	"errors"
)

type SymbolE struct {
}

func (s SymbolE) IsFinal() bool {
	return true
}

func (s SymbolE) Next(n rune) (State, error) {
	switch n {
	case 'e':
		return SymbolE{}, nil
	default:
		return nil, errors.New("incorrect syntax")
	}
}
