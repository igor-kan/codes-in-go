package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9502 evaluates chebyshev collocation node order 9502.
func ComputeChebyshevColloc9502(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
