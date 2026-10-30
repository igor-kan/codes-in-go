package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9252 evaluates chebyshev collocation node order 9252.
func ComputeChebyshevColloc9252(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
