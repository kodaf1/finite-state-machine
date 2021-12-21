package states

type SymbolB struct {
}

func (s SymbolB) IsFinal() bool {
	return false
}

func (s SymbolB) Next(n rune) State {
	switch n {
	case 'b':
		return SymbolB{}
	case 'c':
		return SymbolC{}
	default:
		return Error{}
	}
}
