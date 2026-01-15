package main

import (
	"fmt"
	s "strings"
)

func main(){
	// 97 - 122
	var location [26]int

	var abc string = "abcdefghijklmnopqrstuvwxyz"

	var input string
	fmt.Scanln(&input)

	for index, char := range abc {
		location[index] = s.Index(input, string(char))
	}

	for _, val := range location {
		fmt.Printf("%d ", val)
	}
}