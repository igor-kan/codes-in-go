package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9162 evaluates chebyshev collocation node order 9162.
func ComputeChebyshevColloc9162(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
