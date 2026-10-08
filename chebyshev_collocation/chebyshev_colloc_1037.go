package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1037 evaluates chebyshev collocation node order 1037.
func ComputeChebyshevColloc1037(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
