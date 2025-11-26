package entities

type Player interface {
	GetName() string
	GetSymbol() string
	MakeMove(x, y int)
}
