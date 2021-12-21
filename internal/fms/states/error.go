package states

type Error struct {
}

func (s Error) IsFinal() bool {
	return false
}

func (s Error) Next(n rune) State {
	return s
}
