package goarray

import "fmt"

func Makearray() {
	name := [3]string{"meohyun", "k", "m"}
	names := []string{"a", "b", "c", "d", "e"} // 길이가 정해지지 않은 array를 slice라 한다.
	names = append(names, "f", "g")            // append를 사용해서 slice에 원소를 추가할 수 있다.
	fmt.Println(name)
	fmt.Println(names)
}
