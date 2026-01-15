package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
func init(){
	sc.Split(bufio.ScanWords)
}
func readInt() int {
	sc.Scan()
	n,_ := strconv.Atoi(sc.Text())
	return n
}

func main(){
	var max [3]int = [3]int{0,1,1}

	for i:=1; i<=9; i++ {
		for j:=1; j<=9; j++ {
			n := readInt()
			if n > max[0] { 
				max[0] = n
				max[1] = i
				max[2] = j
			}
		}
	}
	fmt.Printf("%d\n%d %d", max[0], max[1], max[2])
}