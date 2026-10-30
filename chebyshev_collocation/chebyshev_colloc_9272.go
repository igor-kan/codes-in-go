package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9272 evaluates chebyshev collocation node order 9272.
func ComputeChebyshevColloc9272(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
