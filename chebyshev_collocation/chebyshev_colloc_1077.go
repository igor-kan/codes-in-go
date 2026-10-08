package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1077 evaluates chebyshev collocation node order 1077.
func ComputeChebyshevColloc1077(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
