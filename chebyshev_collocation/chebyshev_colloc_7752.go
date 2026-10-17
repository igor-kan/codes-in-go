package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7752 evaluates chebyshev collocation node order 7752.
func ComputeChebyshevColloc7752(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
