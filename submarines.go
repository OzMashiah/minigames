package main

	
import (
	"fmt"
)

func main () {
	PlayerOneBoard := InitiateBoard()
	PlayerTwoBoard := InitiateBoard()
	ShowBoard(PlayerOneBoard)
	ShowBoard(PlayerTwoBoard)
	PlayerOneSubmarines := ChooseSubLoc()
	PlayerTwoSubmarines := ChooseSubLoc()
	fmt.Printf("Player number %d starts.\n", ChooseStarter())
	//while no win -> make move & rotate player 
	CheckWin(PlayerOneBoard, PlayerTwoSubmarines)
	CheckWin(PlayerTwoBoard, PlayerOneSubmarines)

}
