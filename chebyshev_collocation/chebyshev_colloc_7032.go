package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7032 evaluates chebyshev collocation node order 7032.
func ComputeChebyshevColloc7032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
