package chebyshev_collocation

import "math"

// ComputeChebyshevColloc2007 evaluates chebyshev collocation node order 2007.
func ComputeChebyshevColloc2007(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
