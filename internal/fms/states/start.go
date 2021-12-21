package states

type Start struct {
}

func (s Start) IsFinal() bool {
	return false
}

func (s Start) Next(n rune) State {
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
