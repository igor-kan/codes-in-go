package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8517 evaluates chebyshev collocation node order 8517.
func ComputeChebyshevColloc8517(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
