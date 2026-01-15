package main

import(
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var input [100]string

	for i, _ := range input {
		input[i], _ = scanner.ReadString('\n')
	}

	for _, v := range input {
		if v == "\n" {
			break
		}
		fmt.Printf("%s", v)
	}
}