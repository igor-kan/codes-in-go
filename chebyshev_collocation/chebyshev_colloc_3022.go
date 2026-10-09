package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3022 evaluates chebyshev collocation node order 3022.
func ComputeChebyshevColloc3022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
