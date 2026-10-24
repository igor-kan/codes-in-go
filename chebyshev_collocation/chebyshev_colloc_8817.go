package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8817 evaluates chebyshev collocation node order 8817.
func ComputeChebyshevColloc8817(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
