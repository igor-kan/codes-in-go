package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4132 evaluates chebyshev collocation node order 4132.
func ComputeChebyshevColloc4132(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
