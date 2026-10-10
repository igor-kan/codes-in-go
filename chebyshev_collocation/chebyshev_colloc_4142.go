package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4142 evaluates chebyshev collocation node order 4142.
func ComputeChebyshevColloc4142(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
