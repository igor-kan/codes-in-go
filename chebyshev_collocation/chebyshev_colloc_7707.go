package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7707 evaluates chebyshev collocation node order 7707.
func ComputeChebyshevColloc7707(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
