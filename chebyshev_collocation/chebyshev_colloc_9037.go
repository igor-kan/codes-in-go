package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9037 evaluates chebyshev collocation node order 9037.
func ComputeChebyshevColloc9037(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
