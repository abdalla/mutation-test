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

func Divide(a, b int) int {
	if b == 0 {
		return 0
	}

	return a / b
}
