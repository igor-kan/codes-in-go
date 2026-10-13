package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7067 evaluates chebyshev collocation node order 7067.
func ComputeChebyshevColloc7067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
