package calc

func pow(a, b int) int {
	res := 1
	for i := 1; i <= b; i++ {
		res = res * a
	}
	return res
}
