package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4077 evaluates chebyshev collocation node order 4077.
func ComputeChebyshevColloc4077(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
