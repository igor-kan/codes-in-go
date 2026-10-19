package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8142 evaluates chebyshev collocation node order 8142.
func ComputeChebyshevColloc8142(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
