package fuzzing_test

import (
	"testing"

	"github.com/mutation-test/math"
)

func FuzzSum(f *testing.F) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{
			a: 1,
			b: 2,
		},
		{
			a: 0,
			b: 1,
		},
		{
			a: 1,
			b: 0,
		},
		{
			a: 0,
			b: 0,
		},
		{
			a: 1,
			b: -1,
		},
		{
			a: -1,
			b: 1,
		},
		{
			a: 0,
			b: -1,
		},
		{
			a: -1,
			b: 0,
		},
	}

	for _, data := range tests {
		f.Add(data.a, data.b)
	}

	f.Fuzz(func(t *testing.T, a float64, b float64) {
		_ = math.Sum(a, b)

		// if result != (a + b) {
		// 	t.Errorf("sum(%f, %f) = %f; want %f", a, b, result, (a + b))
		// }
	})
}

func FuzzSubtractNeverNegative(f *testing.F) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{
			a: 1,
			b: 2,
		},
		{
			a: 0,
			b: 1,
		},
		{
			a: 1,
			b: 0,
		},
		{
			a: 0,
			b: 0,
		},
		{
			a: 1,
			b: -1,
		},
		{
			a: -1,
			b: 1,
		},
		{
			a: 0,
			b: -1,
		},
		{
			a: -1,
			b: 0,
		},
	}

	for _, data := range tests {
		f.Add(data.a, data.b)
	}

	f.Fuzz(func(t *testing.T, a float64, b float64) {
		_ = math.SubtractNeverNegative(a, b)

		// if result < 0 && result != a {
		// 	t.Errorf("subtractNeverNegative(%f, %f) = %f;", a, b, result)
		// }
	})
}
