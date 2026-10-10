package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4117 evaluates chebyshev collocation node order 4117.
func ComputeChebyshevColloc4117(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
