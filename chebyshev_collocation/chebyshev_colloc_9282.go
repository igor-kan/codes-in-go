package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9282 evaluates chebyshev collocation node order 9282.
func ComputeChebyshevColloc9282(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
