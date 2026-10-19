package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8107 evaluates chebyshev collocation node order 8107.
func ComputeChebyshevColloc8107(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
