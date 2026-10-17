package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7727 evaluates chebyshev collocation node order 7727.
func ComputeChebyshevColloc7727(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
