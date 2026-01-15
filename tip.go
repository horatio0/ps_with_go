// for 문 range
for index, value := range variable {
	// code
}

char - '0' // char to int 형변환

import s "strings" // s.Index(value, string(char)) 없으면 -1 반환

s.Contains(value, string(char)) // 포함 여부 확인, bool 반환
s.Index(value, string(char)) // 포함된 위치 반환, 없으면 -1 반환
s.TrimSpace(value) // 공백, 엔터 제거
s.Split(value, " ") // 특정 문자 기준으로 문자열 나누기, []string 반환

//    병신같은 입력 함수 정상화하는법   \\
var sc = bufio.NewScanner(os.Stdin)

func init(){
	sc.Split(bufio.ScanWords)
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
//     위 코드를 import 아래에 작성     \\