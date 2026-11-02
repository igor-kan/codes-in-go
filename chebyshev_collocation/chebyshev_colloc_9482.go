package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9482 evaluates chebyshev collocation node order 9482.
func ComputeChebyshevColloc9482(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
