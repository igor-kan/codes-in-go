package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9277 evaluates chebyshev collocation node order 9277.
func ComputeChebyshevColloc9277(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
