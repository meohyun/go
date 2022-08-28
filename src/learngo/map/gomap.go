package gomap

import "fmt"

func Makemap() {

	daehyun := map[string]string{"name": "daehyun", "age": "25"}
	fmt.Println(daehyun)

	for key, value := range daehyun {
		fmt.Println(key, value)
	}

}
