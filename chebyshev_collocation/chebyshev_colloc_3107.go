package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3107 evaluates chebyshev collocation node order 3107.
func ComputeChebyshevColloc3107(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
