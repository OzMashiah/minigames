package main

import "fmt"

func ShiftCharacter (c rune, shift int) string {
        shifted := (int(c - 'A') + shift) % 26
        return string('A' + shifted)
}

func main () {	
	fmt.Println(ShiftCharacter('D', 3))
}
