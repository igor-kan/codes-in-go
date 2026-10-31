package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9382 evaluates chebyshev collocation node order 9382.
func ComputeChebyshevColloc9382(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
