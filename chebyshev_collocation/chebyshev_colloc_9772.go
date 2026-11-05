package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9772 evaluates chebyshev collocation node order 9772.
func ComputeChebyshevColloc9772(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
