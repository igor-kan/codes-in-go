package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8942 evaluates chebyshev collocation node order 8942.
func ComputeChebyshevColloc8942(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
