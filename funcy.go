package main

import (
	"math/rand/v2"
	"fmt"
	"strconv"
)

type Submarines struct {
	// The submarines struct that contains 3 submarines and their locations. Sizes of 4, 3, and 2.
	sub4 []string
	sub3 []string
	sub2 []string
}

func InitiateBoard() [11][11]string {
	// This function initialize an empty 10x10 board with column and row indexes, returning a 11x11 board.
	firstcol := [11]string{"@", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J"} // The first column contains the indexes for each row.
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
	board[0] = [11]string{"@", "1", "2", "3", "4" ,"5", "6", "7", "8", "9", "10"} // The first row contains the indexes for each column.
	return board
}

func ShowBoard(board [11][11]string) {
	// This functions gets a board and prints it.
	for i := 0; i < 11; i++ {
        	for j := 0; j < 11; j++ {
            		fmt.Printf("%s ", board[i][j])
        	}
        	fmt.Println() 
    	}
	fmt.Println()
}

func ChooseStarter() int {
	// This function randomly chooses a starter and returns the starting player - 1 or 2.
	return rand.IntN(2)+1
}

func ChooseSubLoc () Submarines {
	// This function creates the submarines struct which contains all submarines locations and returns them.
	submarines := Submarines{
		sub4: GetSubLoc(4),
		sub3: GetSubLoc(3),
		sub2: GetSubLoc(2),
	}
	fmt.Println()

	return submarines
}

func OutOfBoundsLoc (loc string) bool {
	// This function checks for a choosen location (either a move or a submarine location) if its valid.
	// A valid location is 2 or 3 length, first character is between A and J, the rest of the location is between 1 and 10, for example: A10, G5.
	// Returns true if valid, false if not.
	if len(loc) == 2 || len(loc) == 3 {
		if loc[0] >= 'A' && loc[0] <= 'J' {
			num, _ := strconv.Atoi(string(loc[1:]))
			if num >= 1 && num <= 10 {
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
	// This function checks that the chosing submarines length matches the prompt, a 4-length submarine need to be 4 length horizontally or vertically.
	// Return true if the length is correct, false if not.
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
	// This function gets a letter and an integer and returns the letter that follows the input letter by integer amount of steps.
	// For example, input is G, 2 and the function will return I.
	shifted := (int(c - 'A') + shift) % 26
	return string('A' + shifted)
}

func GetSubLoc (length int) []string {
	// This function takes a length of the needed submarine and asks the user to choose the location.
	// returns an array of all the locations. For example, a 3-lenght submarine choosen between G7 and E7 returns [E7 F7 G7].
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
	// This function gets start location, end location and the length to generate the array for the GetSubLoc function.
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

func CheckWin (board [11][11]string, submarines Submarines) bool {
	// This function is being called after every turn to check if we have a winner.
	// Returns true if we have a winner, false if we don't.
	score := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "X" {
				score++
			}
		}
	}
	// Since the sizes of the submarines is fixed (4, 3, and 2) the amount of hits is fixed as well, 9.
	if score == 9 {
		return true
	} else {
		return false
	}
}

func MakeMove (board [11][11]string, submarines Submarines, current_player int) [11][11]string {
	// This function gets a board, a submarines struct, and the current player.
	// Asks the user for a move, makes the move, and returns the new board.
	var move string
	fmt.Printf("Player %d, please choose your move:\n", current_player)
	fmt.Scanln(&move)
	for !OutOfBoundsLoc(move) {
		fmt.Println("Make sure the move is valid:")
		fmt.Scanln(&move)
	}
	submarines_combined := append(append(submarines.sub4[:], submarines.sub3[:]...), submarines.sub2[:]...)
	row, col := TranslateLocToBoard(move)
	for _, loc := range submarines_combined {
		if loc == move {
			// X resembles an hit.
			fmt.Println("It's a hit!")
			board[row][col] = "X"
			return board
		}
	}
	// - resembles a miss.
	fmt.Println("It's a miss.")
	board[row][col] = "-"
	return board
}

func TranslateLocToBoard (loc string) (int, int) {
	// This function gets a locations for example G10 and translates it to row and column, G -> 7, 10 -> 10.
	rowChar := loc[0]
	colNum := loc[1:]

	row := int(rowChar - 'A' + 1) // A -> 1, G -> 7, J -> 10
	col, _ := strconv.Atoi(colNum)

	return row, col
}

func RotatePlayer (current int) int {
	// This function rotates the current player after making a move.
	// Gets the current player and returns the next player.
	if current == 1 {
		return 2
	} else {
		return 1
	}
}

