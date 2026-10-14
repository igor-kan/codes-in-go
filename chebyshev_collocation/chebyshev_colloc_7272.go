package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7272 evaluates chebyshev collocation node order 7272.
func ComputeChebyshevColloc7272(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
