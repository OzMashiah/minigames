package main

	
import (
	"fmt"
)

func main () {
	fmt.Println("Player number", ChooseStarter(), "starts.")
	PlayerOneBoard := InitiateBoard()
	PlayerTwoBoard := InitiateBoard()
	ShowBoard(PlayerOneBoard)
}
