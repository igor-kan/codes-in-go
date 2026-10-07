package chebyshev_collocation

import "math"

// ComputeChebyshevColloc577 evaluates chebyshev collocation node order 577.
func ComputeChebyshevColloc577(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
