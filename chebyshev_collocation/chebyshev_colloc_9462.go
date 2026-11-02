package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9462 evaluates chebyshev collocation node order 9462.
func ComputeChebyshevColloc9462(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
