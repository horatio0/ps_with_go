package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	count, _ := strconv.Atoi(input)

	var inputs []string
	for i := 0; i < count; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)라이언트 오류(ClientCatchedError)
메세지:

An error occurred while processing your request.
		inputs = append(inputs, line)
	}

	var sum int = 0
	for _, line := range inputs {
		sum++
		var m map[rune]int
		m = make(map[rune]int)
		for index, char := range line {
			if m[char] == 0 {
				m[char]++
			} else if m[char] == 1 && rune(line[index-1]) != char {
				sum--
				break
			}
		}		
	}

	fmt.Println(sum)
}