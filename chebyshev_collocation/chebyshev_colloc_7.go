package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7 evaluates chebyshev collocation node order 7.
func ComputeChebyshevColloc7(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
