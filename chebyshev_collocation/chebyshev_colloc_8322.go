package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8322 evaluates chebyshev collocation node order 8322.
func ComputeChebyshevColloc8322(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
