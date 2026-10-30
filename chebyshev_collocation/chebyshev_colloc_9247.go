package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9247 evaluates chebyshev collocation node order 9247.
func ComputeChebyshevColloc9247(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
