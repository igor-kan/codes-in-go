package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9397 evaluates chebyshev collocation node order 9397.
func ComputeChebyshevColloc9397(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
