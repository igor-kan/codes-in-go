package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9032 evaluates chebyshev collocation node order 9032.
func ComputeChebyshevColloc9032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
