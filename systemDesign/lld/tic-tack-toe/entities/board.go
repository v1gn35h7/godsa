package entities

import "fmt"

type Board struct {
	M, N int
	Grid [][]string
}

func NewBoard(m, n int) *Board {
	b := &Board{
		M:    m,
		N:    n,
		Grid: make([][]string, m),
	}
	for i := 0; i < m; i++ {
		b.Grid[i] = make([]string, n)
	}
	return b
}

func (b *Board) PrintBoard() {
	for _, r := range b.Grid {
		fmt.Println(r)
	}
}

func (b *Board) MakeMove(x, y int, symbol string) bool {
	if x < 0 || x >= b.M || y < 0 || y >= b.N {
		return false
	}

	if b.Grid[x][y] != "" {
		return false
	}

	b.Grid[x][y] = symbol
	return true
}

func (b *Board) IsOver() bool {
	for i := 0; i < b.M; i++ {
		for j := 0; j < b.N; j++ {
			if b.Grid[i][j] == "" {
				return false
			}		
	}
	return true
}


// 
func  (b *Board) GetWinner() (bool, *Player) {
	wonSymbol := ""
	// Check rows
    for i := 0; i < 3; i++ {
        if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
            wonSymbol =  board[i][0]
        }
    }

    // Check columns
    for j := 0; j < 3; j++ {
        if board[0][j] != "" && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
            wonSymbol = board[0][j]
        }
    }

    // Check diagonals
    if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
        wonSymbol = board[0][0]
    }
    if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
        retwonSymbol = board[0][2]
    }
	if wonSymbol != "" {
		for _, player := range players {
			if player.GetSymbol() == wonSymbol {
				return true, player
			}
	}

	return false, nil
}