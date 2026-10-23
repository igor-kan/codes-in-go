package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8707 evaluates chebyshev collocation node order 8707.
func ComputeChebyshevColloc8707(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
