package states

type State interface {
	IsFinal() bool
	Next(rune) (State, error)
}
