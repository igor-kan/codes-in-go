package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7442 evaluates chebyshev collocation node order 7442.
func ComputeChebyshevColloc7442(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
