package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8342 evaluates chebyshev collocation node order 8342.
func ComputeChebyshevColloc8342(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
