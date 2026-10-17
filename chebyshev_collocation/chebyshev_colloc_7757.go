package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7757 evaluates chebyshev collocation node order 7757.
func ComputeChebyshevColloc7757(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
