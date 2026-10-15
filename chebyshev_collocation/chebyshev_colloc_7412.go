package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7412 evaluates chebyshev collocation node order 7412.
func ComputeChebyshevColloc7412(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
