package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7257 evaluates chebyshev collocation node order 7257.
func ComputeChebyshevColloc7257(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
