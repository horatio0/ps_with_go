package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a float64 = 0
	var b float64 = 0
	var m map[string]float64 = make (map[string]float64)
	m["A+"] = 4.5
	m["A0"] = 4.0
	m["B+"] = 3.5
	m["B0"] = 3.0
	m["C+"] = 2.5
	m["C0"] = 2.0
	m["D+"] = 1.5
	m["D0"] = 1.0
	m["F"] = 0.0

	for i:=0; i<20; i++ {
		input, _ := reader.ReadString('\n')
		arr := strings.Split(strings.TrimSpace(input), " ")
		if arr[2] != "P" {
			var t float64 = m[arr[2]]
			s, _ := strconv.ParseFloat(arr[1], 64)
			a += float64(s) * t
			b += s
		}
	} 

	if b == 0 { fmt.Printf("0") } else { fmt.Printf("%f", a/b) }
}