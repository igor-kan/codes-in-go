package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9047 evaluates chebyshev collocation node order 9047.
func ComputeChebyshevColloc9047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
