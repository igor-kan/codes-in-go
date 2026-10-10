package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4017 evaluates chebyshev collocation node order 4017.
func ComputeChebyshevColloc4017(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
