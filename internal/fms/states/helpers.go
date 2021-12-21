package states

func GetStartState() State {
	return Start{}
}

func GetErrorState() State {
	return Error{}
}

func GetAllStartStates() []State {
	return []State{
		Start{},
		SymbolA{},
		SymbolB{},
		SymbolC{},
		SymbolD{},
		SymbolE{},
	}
}

func GetStateName(state State) string {
	switch state {
	case Start{}:
		return "Start"
	case SymbolA{}:
		return "A"
	case SymbolB{}:
		return "B"
	case SymbolC{}:
		return "C"
	case SymbolD{}:
		return "D"
	case SymbolE{}:
		return "E"
	case Error{}:
		return "Syntax Error"
	default:
		return "unknown"
	}
}
