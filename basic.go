package main

import (
	"fmt"
	"bufio"
	"os"
)

type person struct {
	name string
	age int
}

func main(){
	
	// 변수 선언
	var a int = 10
	var msg string = "Hello world!"

	// 출력
	fmt.Println(msg, a)

	// 입력
	fmt.Scanln(&msg)
	fmt.Println(msg)

	// 띄어쓰기 포함 입력 받기
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}

	// 함수
	fmt.Println(add(1,2))

	// if 문
	if 1+2!=4 fmt.Println("겠냐?")

	// for 문
	for i:=1; i<=10 ; i++ fmt.Println("무요")

	// 배열
	var c [3]int

	// 슬라이스
	var d []int
	d = []int{1,2,3}
	d = append(d, 4) // 두 슬라이스 합체

	// map
	var m map[int]string
	m = make(map[int]string)
	m[333] = "apple" // 추가 혹은 갱신
	delete(m, 333)   // 삭제
	
	// struct
	p := person{}

	// go routine
	go add(1,2)
	go add(3,4)
	// 비동기 실행
	wait.Wait() // Go 루틴 끝날때까지 대기!

	// 기타 잡
	// defer f.close // 함수 끝나서 return 하기 전에 실행함
	// panic // 실행중인 함수 즉시 종료 (defer 전부 실행 후)
}

// 함수 및 구조체는 앞글자가 대문자면 public, 소문자면 private
func add(a int, b int) {
	return a+b
}
