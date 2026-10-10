package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4067 evaluates chebyshev collocation node order 4067.
func ComputeChebyshevColloc4067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
