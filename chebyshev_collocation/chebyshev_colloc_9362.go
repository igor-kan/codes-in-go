package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9362 evaluates chebyshev collocation node order 9362.
func ComputeChebyshevColloc9362(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
