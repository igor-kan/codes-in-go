package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9667 evaluates chebyshev collocation node order 9667.
func ComputeChebyshevColloc9667(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
