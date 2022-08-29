package gotype

import (
	"fmt"
	"strings"
)

// 다른 언어와 달리 return 값을 여러개 할 수 있음.
// return 값도 type을 지정해줘야 함.
// function 첫글자를 대문자로 해야 export 가능하다.
func Len_Upper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func LenAndUpper(name string) (namelength int, uppername string) {
	defer fmt.Println("이름 저장 완료") // defer은 function이 끝나고 출력된다.
	namelength = len(name)        // 리턴값이 정해졌기에 :=을 쓰지않고 =을 쓴다.
	uppername = strings.ToUpper(name)
	return // naked return이라 한다.
}

// function에서 ...을 통해 여러 파라미터를 받아서 출력할 수 있음.
func Words(words ...string) {
	fmt.Println(words)
}
