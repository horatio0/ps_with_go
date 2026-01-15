package main

import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	scanner := bufio.NewReader(os.Stdin)
	val, _ := scanner.ReadString('\n')
	fmt.Println(val[0])
}