package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8137 evaluates chebyshev collocation node order 8137.
func ComputeChebyshevColloc8137(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
