package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7907 evaluates chebyshev collocation node order 7907.
func ComputeChebyshevColloc7907(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
