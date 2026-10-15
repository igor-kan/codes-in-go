package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7452 evaluates chebyshev collocation node order 7452.
func ComputeChebyshevColloc7452(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
