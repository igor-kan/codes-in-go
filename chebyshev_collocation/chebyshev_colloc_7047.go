package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7047 evaluates chebyshev collocation node order 7047.
func ComputeChebyshevColloc7047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
