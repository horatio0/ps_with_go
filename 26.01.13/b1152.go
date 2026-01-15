// 입력값 없는데 1 출력하는거 생각하기

package main

import (
	f "fmt"
	"bufio"
	"os"
	s "strings"
)

func main(){
	scanner := bufio.NewReader(os.Stdin)
	val, _ := scanner.ReadString('\n')
	val = s.TrimSpace(val)

	if val == "" {
		f.Println(0)
		return
	}
	arr := s.Split(val, " ")

	f.Println(len(arr))
}