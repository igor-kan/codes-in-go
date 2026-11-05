package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9707 evaluates chebyshev collocation node order 9707.
func ComputeChebyshevColloc9707(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
