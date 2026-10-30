package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9292 evaluates chebyshev collocation node order 9292.
func ComputeChebyshevColloc9292(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
