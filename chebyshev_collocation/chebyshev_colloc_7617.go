package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7617 evaluates chebyshev collocation node order 7617.
func ComputeChebyshevColloc7617(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
