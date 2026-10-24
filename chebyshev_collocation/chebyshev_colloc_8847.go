package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8847 evaluates chebyshev collocation node order 8847.
func ComputeChebyshevColloc8847(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
