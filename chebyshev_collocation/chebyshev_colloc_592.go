package chebyshev_collocation

import "math"

// ComputeChebyshevColloc592 evaluates chebyshev collocation node order 592.
func ComputeChebyshevColloc592(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
