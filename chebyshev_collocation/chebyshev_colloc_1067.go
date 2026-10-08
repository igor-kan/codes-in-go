package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1067 evaluates chebyshev collocation node order 1067.
func ComputeChebyshevColloc1067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
