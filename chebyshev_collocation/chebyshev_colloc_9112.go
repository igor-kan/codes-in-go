package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9112 evaluates chebyshev collocation node order 9112.
func ComputeChebyshevColloc9112(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
