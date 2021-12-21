package states

type SymbolC struct {
}

func (s SymbolC) IsFinal() bool {
	return false
}

func (s SymbolC) Next(n rune) State {
	switch n {
	case 'a':
		return SymbolA{}
	case 'b':
		return SymbolB{}
	case 'c':
		return SymbolC{}
	case 'd':
		return SymbolD{}
	default:
		return Error{}
	}
}
