package pointer

import "fmt"

func Pointers() {
	a := 2
	b := &a     // 변수에 a의 메모리 주소를 할당함.
	*b = 202020 // 메모리 주소로 접근한 a에 202020이라는 변수를 할당함.
	fmt.Println(a)
}
