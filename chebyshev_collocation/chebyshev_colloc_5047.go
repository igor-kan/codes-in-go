package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5047 evaluates chebyshev collocation node order 5047.
func ComputeChebyshevColloc5047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
