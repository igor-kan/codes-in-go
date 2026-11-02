package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9532 evaluates chebyshev collocation node order 9532.
func ComputeChebyshevColloc9532(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
