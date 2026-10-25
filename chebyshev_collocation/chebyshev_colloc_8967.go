package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8967 evaluates chebyshev collocation node order 8967.
func ComputeChebyshevColloc8967(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
