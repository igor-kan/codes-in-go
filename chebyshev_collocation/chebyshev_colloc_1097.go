package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1097 evaluates chebyshev collocation node order 1097.
func ComputeChebyshevColloc1097(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
