package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9677 evaluates chebyshev collocation node order 9677.
func ComputeChebyshevColloc9677(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
