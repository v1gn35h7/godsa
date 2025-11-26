package main

import "github.com/vicky1115050/godsa/systemDesign/lld/tic-tack-toe/entities"

func main() {

	var game entities.Game

	m, n := 3, 3

	var player1, player2 entities.Player

	player1 = entities.NewHumanPlayer("Alice", "X")
	player2 = entities.NewHumanPlayer("Bob", "O")

	board := entities.NewBoard(m, n)

	game = entities.NewBoardGame([]entities.Player{player1, player2}, board)

	game.Start()

}

//
