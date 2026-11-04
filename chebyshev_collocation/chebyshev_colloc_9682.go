package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9682 evaluates chebyshev collocation node order 9682.
func ComputeChebyshevColloc9682(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
