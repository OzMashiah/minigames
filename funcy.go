package main

import (
	"math/rand/v2"
	"fmt"
	"strconv"
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
        	fmt.Println() 
    	}
}

func ChooseStarter() int {
	return rand.IntN(2)+1
}

func ChooseSubLoc submarines () {
	type submarines struct {
		sub4 [4]string
		sub3 [3]string
		sub2 [2]string
	}
	
	var start4 string
	var end4 string
	var start3 string
	var end3 string
	var start2 string
	var end2 string

	fmt.Println("Enter the starting location of the 4-length submarine: ")
    	fmt.Scanln(&start4)
	for !LocCheck(start4) {
		fmt.Println("Location not within board bounds, please try again (examples: A4, J9, B1): ")
		fmt.Scanln(&start4)
	}
	fmt.Println("Enter the ending location of the 4-length submarine, make sure its 4 length: ")
	fmt.Scanln(&end4)
	for !LocCheck(end4) {
		fmt.Println("Location not within board bounds, please try again (examples: A4, J9, B1): ")
		fmt.Scanln(&end4)
	}
	for !CheckSubLen(start4, end4, 4) {
		fmt.Println("Ending location is not 4 length from the starting location, please try again (examples A4-A7, A4-D4): "
		fmt.Scanln(&end4)
	}
	
}

func OutOfBoundsLoc (loc string) bool {
	if loc[0] >= 'A' && loc[0] <= 'J' {
		num, err := strconv.Atoi(string(loc[1]))
		if err == nil && num >= 1 && num <= 10 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func CheckLocLength (loc string) bool {
	if len(loc) == 2 {
		return true
	} else {
		return false
	}
}

func LocCheck (loc string) bool {
	if CheckLocLength(loc) && OutOfBoundsLoc(loc) {
		return true
	} else {
		return false
	}
}

func CheckSubLen (start string, end string, length int) bool {
	if ShiftCharacter(rune(start[0]), length-1) == string(end[0]) || int(start[1]) + length-1 == int(end[1]) {
		return true
	} else {
		return false
	}
}

func ShiftCharacter (c rune, shift int) string {
	shifted := (int(c - 'A') + shift) % 26
	return string('A' + shifted)
}




