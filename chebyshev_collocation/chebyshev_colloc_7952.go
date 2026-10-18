package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7952 evaluates chebyshev collocation node order 7952.
func ComputeChebyshevColloc7952(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
