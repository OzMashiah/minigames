package main

	
import (
	"fmt"
)

func main () {
	// Initialize boards
	PlayerOneBoard := InitiateBoard()
	PlayerTwoBoard := InitiateBoard()

	// Choose submarines locations
	fmt.Println("Player 1, choose your submarines locations")
	PlayerOneSubmarines := ChooseSubLoc()
	fmt.Println("Player 2, choose your submarines locations")
	PlayerTwoSubmarines := ChooseSubLoc()

	// Choose starter
	current_player := ChooseStarter()
	fmt.Printf("Player number %d starts.\n", current_player)

	// Play game.
	for !CheckWin(PlayerOneBoard, PlayerTwoSubmarines) && !CheckWin(PlayerTwoBoard, PlayerOneSubmarines) {
		if current_player == 1 {
			PlayerOneBoard = MakeMove(PlayerOneBoard, PlayerTwoSubmarines, current_player)
			ShowBoard(PlayerOneBoard)
		} else if current_player == 2 {
			PlayerTwoBoard = MakeMove(PlayerTwoBoard, PlayerOneSubmarines, current_player)
			ShowBoard(PlayerTwoBoard)
		}
		current_player = RotatePlayer(current_player)
	}

	// Ending.
	current_player = RotatePlayer(current_player)
	fmt.Printf("Player number %d wins!\n", current_player)
	fmt.Println("The board of player 1:")
	ShowBoard(PlayerOneBoard)
	fmt.Println("The board of player 2:")
	ShowBoard(PlayerTwoBoard)
}
