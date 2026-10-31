package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9377 evaluates chebyshev collocation node order 9377.
func ComputeChebyshevColloc9377(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
