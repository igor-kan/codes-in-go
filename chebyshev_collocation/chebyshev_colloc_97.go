package chebyshev_collocation

import "math"

// ComputeChebyshevColloc97 evaluates chebyshev collocation node order 97.
func ComputeChebyshevColloc97(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
