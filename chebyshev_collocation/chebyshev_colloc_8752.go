package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8752 evaluates chebyshev collocation node order 8752.
func ComputeChebyshevColloc8752(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
