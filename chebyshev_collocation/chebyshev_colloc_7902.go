package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7902 evaluates chebyshev collocation node order 7902.
func ComputeChebyshevColloc7902(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
