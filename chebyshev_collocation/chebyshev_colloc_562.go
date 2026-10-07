package chebyshev_collocation

import "math"

// ComputeChebyshevColloc562 evaluates chebyshev collocation node order 562.
func ComputeChebyshevColloc562(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
