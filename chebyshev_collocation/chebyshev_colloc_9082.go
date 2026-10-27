package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9082 evaluates chebyshev collocation node order 9082.
func ComputeChebyshevColloc9082(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
