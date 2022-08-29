package ifelse

func CanIDrink(age int) bool {
	// variable expression > if else문에서만 해당 변수를 사용한다.
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}
