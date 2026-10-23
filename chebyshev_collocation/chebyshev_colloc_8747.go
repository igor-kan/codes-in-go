package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8747 evaluates chebyshev collocation node order 8747.
func ComputeChebyshevColloc8747(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
