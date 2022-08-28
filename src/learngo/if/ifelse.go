package ifelse

func CanIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}
