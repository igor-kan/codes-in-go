package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7427 evaluates chebyshev collocation node order 7427.
func ComputeChebyshevColloc7427(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
