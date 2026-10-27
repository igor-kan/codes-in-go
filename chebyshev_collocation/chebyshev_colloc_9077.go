package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9077 evaluates chebyshev collocation node order 9077.
func ComputeChebyshevColloc9077(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
