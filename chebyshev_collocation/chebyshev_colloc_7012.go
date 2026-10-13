package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7012 evaluates chebyshev collocation node order 7012.
func ComputeChebyshevColloc7012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
