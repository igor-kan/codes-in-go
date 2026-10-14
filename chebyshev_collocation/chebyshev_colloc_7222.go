package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7222 evaluates chebyshev collocation node order 7222.
func ComputeChebyshevColloc7222(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
