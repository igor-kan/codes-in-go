package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9642 evaluates chebyshev collocation node order 9642.
func ComputeChebyshevColloc9642(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
