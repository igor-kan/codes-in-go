package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9242 evaluates chebyshev collocation node order 9242.
func ComputeChebyshevColloc9242(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
