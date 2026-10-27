package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9127 evaluates chebyshev collocation node order 9127.
func ComputeChebyshevColloc9127(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
