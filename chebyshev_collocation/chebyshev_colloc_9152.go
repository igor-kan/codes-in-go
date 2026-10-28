package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9152 evaluates chebyshev collocation node order 9152.
func ComputeChebyshevColloc9152(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
