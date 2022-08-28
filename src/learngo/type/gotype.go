package gotype

import (
	"fmt"
	"strings"
)

// 다른 언어와 달리 return 값을 여러개 할 수 있음.
// return 값도 type을 지정해줘야 함.
func Len_Upper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func LenAndUpper(name string) (namelength int, uppername string) {
	defer fmt.Println("이름 저장 완료")
	namelength = len(name)
	uppername = strings.ToUpper(name)
	return
}

// function에서 ...을 통해 여러 파라미터를 받아서 출력할 수 있음.
func Words(words ...string) {
	fmt.Println(words)
}
