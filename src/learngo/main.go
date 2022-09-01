package main // 컴파일을 위해서는 main package를 사용해야한다.

import (
	// "fmt"
	// 	ifelse "learngo/if"
	// 	loop "learngo/loop"
	// 	pointer "learngo/pointer"
	// goarray "learngo/array"
	// gomap "learngo/map"
	// gostruct "learngo/struct"
	//accounts "learngo/banking"
	"fmt"
	mydict "learngo/dict"
)

func main() {
	// loop.Loopnumber(7)
	// loop.Loopnumberrange(1, 2, 3, 4, 5, 6, 7)
	// plus := loop.Plusnumbers(1, 2, 3, 4, 5, 6, 7)
	// fmt.Println(plus)
	// fmt.Println(ifelse.CanIDrink(18))
	// pointer.Pointers()
	// goarray.Makearray()
	// gomap.Makemap()
	// gostruct.Makestruct()
	//account := accounts.NewAccount("대현")
	// account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance())
	// account.ChangeOwner("현다이")
	// fmt.Println(account.Owner())
	// fmt.Println(account)
	dictonary := mydict.Dictionary{"First": "first word"}
	//fmt.Println(dictonary)
	// value, err := dictonary.Search("Second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(value)
	// }
	adderr := dictonary.Add("Second", "second word")
	if adderr != nil {
		fmt.Println(adderr)
	}

	// updateerr := dictonary.Update("First", "Second")
	// if updateerr != nil {
	// 	fmt.Println(updateerr)
	// }
	// fmt.Println(dictonary)

	err := dictonary.Worddelete("First")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictonary)
	}

}
