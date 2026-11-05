package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9747 evaluates chebyshev collocation node order 9747.
func ComputeChebyshevColloc9747(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
