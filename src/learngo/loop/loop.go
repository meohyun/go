package loop

import "fmt"

func Loopnumber(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

func Loopnumberrange(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}

func Plusnumbers(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += i
	}
	return total
}
