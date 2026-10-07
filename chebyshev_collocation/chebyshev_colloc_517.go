package chebyshev_collocation

import "math"

// ComputeChebyshevColloc517 evaluates chebyshev collocation node order 517.
func ComputeChebyshevColloc517(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
