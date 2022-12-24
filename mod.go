package main

func Mod(a, b float64) float64 {

	if a > b {
		for a > b {
			a = a - b
		}
	}
	return a
}
