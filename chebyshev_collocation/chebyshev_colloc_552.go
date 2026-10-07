package chebyshev_collocation

import "math"

// ComputeChebyshevColloc552 evaluates chebyshev collocation node order 552.
func ComputeChebyshevColloc552(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
