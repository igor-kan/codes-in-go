package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9452 evaluates chebyshev collocation node order 9452.
func ComputeChebyshevColloc9452(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
