package calc

import "golang.org/x/exp/constraints"

func Sum[T constraints.Float](a, b T) T {
	return a + b
}
