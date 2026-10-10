package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4022 evaluates chebyshev collocation node order 4022.
func ComputeChebyshevColloc4022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
