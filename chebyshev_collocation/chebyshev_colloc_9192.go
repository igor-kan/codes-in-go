package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9192 evaluates chebyshev collocation node order 9192.
func ComputeChebyshevColloc9192(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
