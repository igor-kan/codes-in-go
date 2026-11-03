package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9612 evaluates chebyshev collocation node order 9612.
func ComputeChebyshevColloc9612(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
