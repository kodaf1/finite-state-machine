package states

type SymbolE struct {
}

func (s SymbolE) IsFinal() bool {
	return true
}

func (s SymbolE) Next(n rune) State {
	switch n {
	case 'e':
		return SymbolE{}
	default:
		return Error{}
	}
}
