package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1012 evaluates chebyshev collocation node order 1012.
func ComputeChebyshevColloc1012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
