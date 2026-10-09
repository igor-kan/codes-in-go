package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3052 evaluates chebyshev collocation node order 3052.
func ComputeChebyshevColloc3052(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
