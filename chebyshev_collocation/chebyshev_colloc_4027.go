package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4027 evaluates chebyshev collocation node order 4027.
func ComputeChebyshevColloc4027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
