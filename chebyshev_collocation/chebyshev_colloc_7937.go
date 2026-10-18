package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7937 evaluates chebyshev collocation node order 7937.
func ComputeChebyshevColloc7937(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
