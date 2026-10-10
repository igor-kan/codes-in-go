package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4007 evaluates chebyshev collocation node order 4007.
func ComputeChebyshevColloc4007(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
