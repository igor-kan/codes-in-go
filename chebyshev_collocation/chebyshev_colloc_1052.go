package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1052 evaluates chebyshev collocation node order 1052.
func ComputeChebyshevColloc1052(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
