package main

import (
	"math/rand/v2"
	"fmt"
)

func InitiateBoard() [11][11]string {
	firstcol := [11]string{"@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	var board [11][11]string
	for i := range board {
		for j := range board[i] {
			board[i][j] = "?"
			if j == 0 {
				board[i][j] = firstcol[i]
			} else {
				board[i][j] = "?"
			}
		}
	}
	board[0] = [11]string{"@", "1", "2", "3", "4" ,"5", "6", "7", "8", "9", "10"}
	return board
}

func ShowBoard(board [11][11]string) {
	for i := 0; i < 11; i++ {
        	for j := 0; j < 11; j++ {
            		fmt.Printf("%s ", board[i][j])
        	}
        	fmt.Println() // New line after each row
    	}
}

func ChooseStarter() int {
	return rand.IntN(2)+1
}

