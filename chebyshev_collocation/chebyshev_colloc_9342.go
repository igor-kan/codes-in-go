package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9342 evaluates chebyshev collocation node order 9342.
func ComputeChebyshevColloc9342(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
