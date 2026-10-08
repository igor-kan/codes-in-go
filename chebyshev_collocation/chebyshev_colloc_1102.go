package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1102 evaluates chebyshev collocation node order 1102.
func ComputeChebyshevColloc1102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
