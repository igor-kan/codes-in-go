package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9552 evaluates chebyshev collocation node order 9552.
func ComputeChebyshevColloc9552(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
