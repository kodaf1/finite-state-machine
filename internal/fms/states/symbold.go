package states

type SymbolD struct {
}

func (s SymbolD) IsFinal() bool {
	return true
}

func (s SymbolD) Next(n rune) State {
	switch n {
	case 'e':
		return SymbolE{}
	default:
		return Error{}
	}
}
