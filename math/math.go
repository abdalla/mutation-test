package math

func Sum(a, b float64) float64 {
	return a + b
}

func SubtractNeverNegative(a, b float64) float64 {
	if a >= b {
		return a - b
	}
	return a
}
