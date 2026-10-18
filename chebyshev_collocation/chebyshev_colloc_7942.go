package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7942 evaluates chebyshev collocation node order 7942.
func ComputeChebyshevColloc7942(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
