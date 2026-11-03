package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9617 evaluates chebyshev collocation node order 9617.
func ComputeChebyshevColloc9617(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
