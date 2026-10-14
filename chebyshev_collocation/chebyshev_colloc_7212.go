package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7212 evaluates chebyshev collocation node order 7212.
func ComputeChebyshevColloc7212(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
