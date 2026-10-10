package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4052 evaluates chebyshev collocation node order 4052.
func ComputeChebyshevColloc4052(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
