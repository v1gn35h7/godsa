package entities

type HumanPlayer struct {
	Name   string
	Symbol string
}

func NewHumanPlayer(name, symbol string) *HumanPlayer {
	return &HumanPlayer{Name: name, Symbol: symbol}
}

func (hp *HumanPlayer) MakeMove(x, y int) {
	// Implementation for human player making a move
}

func (hp *HumanPlayer) GetName() string {
	return hp.Name
}

func (hp *HumanPlayer) GetSymbol() string {
	return hp.Symbol
}
