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
	if len(loc) == 2 || len(loc) == 3 {
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
	} else {
		return false
	}
}

func CheckSubLen (start string, end string, length int) bool {
	start_num_int, _ := strconv.Atoi(start[1:])
	end_num_int, _ := strconv.Atoi(end[1:])
	if ShiftCharacter(rune(start[0]), length-1) == string(end[0]) || start_num_int + length - 1 == end_num_int ||
	string(start[0]) == ShiftCharacter(rune(end[0]), length-1) || start_num_int == end_num_int + length - 1 {
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
	start_char := start[:1]
	end_char := end[:1]
	start_num, _ := strconv.Atoi(start[1:])
	end_num, _ := strconv.Atoi(end[1:])

	if start_char == end_char {
		// horizontal
		if start_num < end_num {
			// foreward
			for i := 0; i < length; i++ {
				str_num := strconv.Itoa(start_num+i)
				sub[i] = string(start_char) + str_num
			}
		} else {
			// backward
			for i := 0; i < length; i++ {
				str_num := strconv.Itoa(start_num-i)
				sub[i] = string(start_char) + str_num
			}
		}
	} else {
		// vertical
		start_num_str := strconv.Itoa(start_num)
		end_num_str := strconv.Itoa(end_num)
		if start_char < end_char {
			// foreward
			for i := 0; i < length; i++ {
				sub[i] = ShiftCharacter(rune(start_char[0]), i) + end_num_str
			}
		} else {
			// backward
			for i := 0; i < length; i++ {
				sub[i] = ShiftCharacter(rune(end_char[0]), i) + start_num_str
			}
		}
	}
	return sub
}

func CheckWin (board [11][11]string, submarines Submarines) {
	score := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 9 {
				return true
		}
	}
	return false
}

func MakeMove (board [11][11]string, submarines Submarines, move string) {
	submarines_combined := append(sub4[:], sub3[:], sub2[:])
	row, col := TranslateLocToBoard(move)
	for _, loc := range submarines_combined {
		if loc == move {
			// X resembles an hit.
			board[row][col] = 'X'
		} else {
			// - resembles a miss.
			board[row][col] = '-'
		}

}

func TranslateLocToBoard (loc string) (int, int) {
	rowChar := loc[0]
	colNum := loc[1:]

	row := int(rowChar - 'A' + 1) // A -> 1, G -> 7, J -> 10
	col, _ := strconv.Atoi(colNum)

	return row, col
}

