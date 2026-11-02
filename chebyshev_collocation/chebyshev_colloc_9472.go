package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9472 evaluates chebyshev collocation node order 9472.
func ComputeChebyshevColloc9472(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
