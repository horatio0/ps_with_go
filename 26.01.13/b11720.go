package main

import "fmt"

func main(){
	var count int
	fmt.Scanln(&count)

	var input string
	fmt.Scanln(&input)

	var sum int64 = 0
	for _, char := range input {
		sum += int64(char - '0')
	}
	fmt.Println(sum)
}