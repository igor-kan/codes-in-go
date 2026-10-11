package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5097 evaluates chebyshev collocation node order 5097.
func ComputeChebyshevColloc5097(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
