package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8102 evaluates chebyshev collocation node order 8102.
func ComputeChebyshevColloc8102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
