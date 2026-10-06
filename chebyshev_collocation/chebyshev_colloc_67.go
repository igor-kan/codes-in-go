package chebyshev_collocation

import "math"

// ComputeChebyshevColloc67 evaluates chebyshev collocation node order 67.
func ComputeChebyshevColloc67(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
