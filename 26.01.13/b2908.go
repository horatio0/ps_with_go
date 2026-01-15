package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

func main(){
	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	input = strings.TrimSpace(input)
	

	arr := strings.Split(input, " ")
	var result [2]int

	for index, v := range arr {
		runes := []rune(v)
		for i := 0; i < len(runes)/2; i++ {
			runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
		}
		result[index],_ = strconv.Atoi(string(runes))
	}

	if result[0] > result[1] {
		fmt.Println(result[0])
	} else {
		fmt.Println(result[1])
	}
}