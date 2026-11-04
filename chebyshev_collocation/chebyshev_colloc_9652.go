package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9652 evaluates chebyshev collocation node order 9652.
func ComputeChebyshevColloc9652(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
