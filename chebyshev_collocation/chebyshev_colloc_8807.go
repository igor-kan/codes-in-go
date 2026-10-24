package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8807 evaluates chebyshev collocation node order 8807.
func ComputeChebyshevColloc8807(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
