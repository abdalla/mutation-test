package math

func sum(a, b float64) float64 {
	return a + b
}

func subtractNeverNegative(a, b float64) float64 {
	if a >= b {
		return a - b
	}
	return a
}
