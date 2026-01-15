package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var t int
	fmt.Scanln(&t)
	scanner := bufio.NewReader(os.Stdin)
	var input []string
	
	for i := 0;i < t; i++ {
		val, _ := scanner.ReadString('\n')
		input = append(input, string(val))
	}

	for i := 0; i < t; i++ {
		val := []rune(input[i])
		fmt.Printf("%s%s\n", string(val[0]), string(val[len(val)-2]))
	}
}