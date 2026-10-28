package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9202 evaluates chebyshev collocation node order 9202.
func ComputeChebyshevColloc9202(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
