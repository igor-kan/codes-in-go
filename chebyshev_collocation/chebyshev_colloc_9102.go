package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9102 evaluates chebyshev collocation node order 9102.
func ComputeChebyshevColloc9102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
