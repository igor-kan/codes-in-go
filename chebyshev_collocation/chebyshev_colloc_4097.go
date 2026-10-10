package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4097 evaluates chebyshev collocation node order 4097.
func ComputeChebyshevColloc4097(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
