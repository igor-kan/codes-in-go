package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7052 evaluates chebyshev collocation node order 7052.
func ComputeChebyshevColloc7052(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
