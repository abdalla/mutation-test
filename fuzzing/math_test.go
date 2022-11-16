package fuzzing_test

import (
	"testing"

	"github.com/mutation-test/math"
)

func FuzzDivide(f *testing.F) {
	tests := []struct {
		a int
		b int
	}{
		{
			a: 1,
			b: 2,
		},
	}

	for _, data := range tests {
		f.Add(data.a, data.b)
	}

	f.Fuzz(func(t *testing.T, a int, b int) {
		_ = math.Divide(a, b)
	})
}
