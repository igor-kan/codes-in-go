package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4127 evaluates chebyshev collocation node order 4127.
func ComputeChebyshevColloc4127(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
