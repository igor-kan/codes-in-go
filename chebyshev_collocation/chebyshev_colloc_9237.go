package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9237 evaluates chebyshev collocation node order 9237.
func ComputeChebyshevColloc9237(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
