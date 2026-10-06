package chebyshev_collocation

import "math"

// ComputeChebyshevColloc107 evaluates chebyshev collocation node order 107.
func ComputeChebyshevColloc107(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
