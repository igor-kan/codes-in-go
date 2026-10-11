package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5037 evaluates chebyshev collocation node order 5037.
func ComputeChebyshevColloc5037(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
