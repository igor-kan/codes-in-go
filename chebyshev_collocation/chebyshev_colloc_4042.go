package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4042 evaluates chebyshev collocation node order 4042.
func ComputeChebyshevColloc4042(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
