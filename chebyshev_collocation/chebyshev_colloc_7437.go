package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7437 evaluates chebyshev collocation node order 7437.
func ComputeChebyshevColloc7437(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
