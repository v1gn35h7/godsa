package entities

type BoardGameStrategy interface {
	CheckWin(board *Board, symbol string) bool
	MakeMove(x, y int, symbol string) bool
}
