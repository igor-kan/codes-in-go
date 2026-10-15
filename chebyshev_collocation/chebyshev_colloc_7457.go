package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7457 evaluates chebyshev collocation node order 7457.
func ComputeChebyshevColloc7457(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
