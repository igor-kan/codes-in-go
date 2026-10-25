package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8887 evaluates chebyshev collocation node order 8887.
func ComputeChebyshevColloc8887(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
