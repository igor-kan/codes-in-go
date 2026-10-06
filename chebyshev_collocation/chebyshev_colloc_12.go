package chebyshev_collocation

import "math"

// ComputeChebyshevColloc12 evaluates chebyshev collocation node order 12.
func ComputeChebyshevColloc12(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
