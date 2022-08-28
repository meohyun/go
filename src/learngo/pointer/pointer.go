package pointer

import "fmt"

func Pointers() {
	a := 2
	b := &a
	*b = 202020
	fmt.Println(a)
}
