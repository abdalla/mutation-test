package math

import "testing"

func TestSum(t *testing.T) {
	a := 8.5
	b := 1.5

	result := sum(a, b)

	if result != 10.0 {
		t.Errorf("sum(8.5, 1.5) = %f; want 10", result)
	}
}

func TestSubtractNeverNegative(t *testing.T) {
	a := 9.4
	b := 9.3

	result := subtractNeverNegative(a, b)
	if result < 0 {
		t.Error("Non Negative")
	}

	result = subtractNeverNegative(a, 20.3)
	if result != a {
		t.Error("Errorrrrr")
	}

	// Although you put the >= in the line 9 on math.go file,
	// the mutation test still failing cuz u are not testing it.
	// so, uncomment the code bellow
	// result = subtractNeverNegative(a, a)
	// if result != 0 {
	// 	t.Error("Errorrrrr")
	// }
}
