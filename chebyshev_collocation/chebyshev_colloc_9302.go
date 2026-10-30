package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9302 evaluates chebyshev collocation node order 9302.
func ComputeChebyshevColloc9302(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
