package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3037 evaluates chebyshev collocation node order 3037.
func ComputeChebyshevColloc3037(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
