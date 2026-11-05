package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9727 evaluates chebyshev collocation node order 9727.
func ComputeChebyshevColloc9727(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
