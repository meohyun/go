package gostruct

import "fmt"

// struct에 들어갈 변수들의 type을 정해준다.
type person struct {
	name  string
	age   int
	hobby []string
}

func Makestruct() {
	hobby := []string{"baseball", "game"}
	daehyun := person{name: "daehyun", age: 25, hobby: hobby}
	daehyun.name = "meohyun"
	fmt.Println(daehyun)
	fmt.Println(daehyun.name)
}
