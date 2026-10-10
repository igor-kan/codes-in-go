package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4032 evaluates chebyshev collocation node order 4032.
func ComputeChebyshevColloc4032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
