package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9322 evaluates chebyshev collocation node order 9322.
func ComputeChebyshevColloc9322(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
