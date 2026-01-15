package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var m map[string]int
	m = make(map[string]int)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToUpper(input)
	for _, char := range input {
		m[string(char)]++
	}

	var result int = 0
	var check int = 0
	var a string
	for key, count := range m {
		if result < count { a = key; result = count } else if result == count { check = count }
	}
	if check == result { fmt.Println("?") } else {fmt.Println(a)}
}