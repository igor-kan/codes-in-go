package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9587 evaluates chebyshev collocation node order 9587.
func ComputeChebyshevColloc9587(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
