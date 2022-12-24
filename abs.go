package calc

func Abs(a int) int {
	if a < 0 {
		a *= -1
		return a
	}
	return a
}
