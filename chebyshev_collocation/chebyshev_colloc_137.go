package chebyshev_collocation

import "math"

// ComputeChebyshevColloc137 evaluates chebyshev collocation node order 137.
func ComputeChebyshevColloc137(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
