package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9332 evaluates chebyshev collocation node order 9332.
func ComputeChebyshevColloc9332(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
