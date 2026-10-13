package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7037 evaluates chebyshev collocation node order 7037.
func ComputeChebyshevColloc7037(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
