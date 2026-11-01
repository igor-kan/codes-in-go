package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9387 evaluates chebyshev collocation node order 9387.
func ComputeChebyshevColloc9387(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
