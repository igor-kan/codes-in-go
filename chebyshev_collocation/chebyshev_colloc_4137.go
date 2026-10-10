package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4137 evaluates chebyshev collocation node order 4137.
func ComputeChebyshevColloc4137(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
