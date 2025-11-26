package entities

import "fmt"

type BoardGame struct {
	Players       []Player
	Board         *Board
	CurrentPlayer *Player
}

func NewBoardGame(players []Player, board *Board) *BoardGame {
	return &BoardGame{
		Players: players,
		Board:   board,
	}
}

func (bg *BoardGame) Start() {
	// Implementation for starting the board game
	moves := [][]int{[]int{0, 0}, []int{1, 1}, []int{0, 1}, []int{1, 0}, []int{0, 2}}

	for _, move := range moves {
		x, y := move[0], move[1]

		s := bg.MakeMove(x, y)

		if s {
			bg.Board.PrintBoard()

			over := bg.Board.IsOver()
			if over {
				fmt.Println("Game Over!")
				bg.PrintWinner()
			}
		}

	}
}

func (bg *BoardGame) GetCurrentPlayer() *Player {
	return bg.CurrentPlayer
}

func (bg *BoardGame) MakeMove(x, y int, symbol string) bool {
	sucess := bg.Board.MakeMove(x, y, symbol)
	if !sucess {
		fmt.Println("Invalid move")
		return false
	}

	return true
}

func (bg *BoardGame) PrintWinner() {
	won, player := bg.Board.GetWinner()
	if won {
		fmt.Printf("Winner is %s\n", player.GetName())
	} else {
		fmt.Println("It's a draw!")
	}
}
