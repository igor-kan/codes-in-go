package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3007 evaluates chebyshev collocation node order 3007.
func ComputeChebyshevColloc3007(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
