package states

type SymbolA struct {
}

func (s SymbolA) IsFinal() bool {
	return false
}

func (s SymbolA) Next(n rune) State {
	switch n {
	case 'c':
		return SymbolC{}
	default:
		return Error{}
	}
}
