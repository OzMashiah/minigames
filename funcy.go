package main

import (
	"math/rand/v2"
	"fmt"
	"strconv"
)

type Submarines struct {
	sub4 []string
	sub3 []string
	sub2 []string
}

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
	fmt.Println()
}

func ChooseStarter() int {
	return rand.IntN(2)+1
}

func ChooseSubLoc () Submarines {
	submarines := Submarines{
		sub4: GetSubLoc(4),
		sub3: GetSubLoc(3),
		sub2: GetSubLoc(2),
	}

	return submarines
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

func GetSubLoc (length int) []string {
	var start string
        var end string

        fmt.Printf("Enter the starting location of the %d-length submarine:\n", length)
        fmt.Scanln(&start)
        for !OutOfBoundsLoc(start) {
                fmt.Println("Location not within board bounds, please try again (examples: A4, J9, B1): ")
                fmt.Scanln(&start)
        }
        fmt.Printf("Enter the ending location of the %d-length submarine, make sure its %d length:\n", length, length)
        fmt.Scanln(&end)
        for !OutOfBoundsLoc(end) {
                fmt.Println("Location not within board bounds, please try again (examples: A4, J9, B1): ")
                fmt.Scanln(&end)
        }
        for !CheckSubLen(start, end, length) {
		fmt.Printf("Ending location is not %d length from the starting location\n" + 
		"Examples: 4-length - A4-A7 / A4-D4, 3-length - B2-B4 / B2-D2, 2-length - G9-G10 / G9-H9:\n", length)
                fmt.Scanln(&end)
        }
	return GenerateSub(start, end, length)
}

func GenerateSub (start string, end string, length int) []string {
	var sub []string
	sub = make([]string, length)
	if start[0] == end[0] {
		// horizontal
		for i := 0; i < length; i++ {
			sub[i] = string(start[0]) + string(int(start[1])+i)
		}
	} else {
		// vertical
		for i := 0; i < length; i++ {
			sub[i] = ShiftCharacter(rune(start[0]), i) + string(start[1])
		}
	}
	return sub
}



