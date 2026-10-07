package chebyshev_collocation

import "math"

// ComputeChebyshevColloc607 evaluates chebyshev collocation node order 607.
func ComputeChebyshevColloc607(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
