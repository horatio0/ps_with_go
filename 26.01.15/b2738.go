package main

import (
	"fmt"
	"os"
	//"strings"
	"bufio"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init(){
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main(){
	
	N := nextInt()
	M := nextInt()
	
	var a [100][100]int
	var b [100][100]int
	var i,j int

	for i=0; i<N; i++ {
		for j=0; j<M; j++ {
			a[i][j] = nextInt()
		}
	}
	for i=0; i<N; i++ {
		for j=0; j<M; j++ {
			b[i][j] = nextInt()
		}
	}
	for i=0; i<N; i++ {
		for j=0; j<M; j++ {
			fmt.Printf("%d ", a[i][j]+b[i][j])
		}
		fmt.Println()
	}
}