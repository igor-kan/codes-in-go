package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9227 evaluates chebyshev collocation node order 9227.
func ComputeChebyshevColloc9227(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
