package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1032 evaluates chebyshev collocation node order 1032.
func ComputeChebyshevColloc1032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
