package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9367 evaluates chebyshev collocation node order 9367.
func ComputeChebyshevColloc9367(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
