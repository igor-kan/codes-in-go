package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8357 evaluates chebyshev collocation node order 8357.
func ComputeChebyshevColloc8357(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
