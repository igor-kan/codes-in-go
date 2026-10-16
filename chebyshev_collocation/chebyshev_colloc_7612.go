package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7612 evaluates chebyshev collocation node order 7612.
func ComputeChebyshevColloc7612(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
