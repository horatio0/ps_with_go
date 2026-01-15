package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var m map[string]int
	m = make(map[string]int)

	m["ABC"] = 3
	m["DEF"] = 4
	m["GHI"] = 5
	m["JKL"] = 6
	m["MNO"] = 7
	m["PQRS"] = 8
	m["TUV"] = 9
	m["WXYZ"] = 10

	s, _ := scanner.ReadString('\n')
	var result int = 0

	for _, char := range s {
		for key, val := range m {
			if strings.Contains(key, string(char)) {
				result += val
				break
			}
		}
	}

	fmt.Println(result)
}