package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7252 evaluates chebyshev collocation node order 7252.
func ComputeChebyshevColloc7252(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
