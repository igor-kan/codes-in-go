package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9562 evaluates chebyshev collocation node order 9562.
func ComputeChebyshevColloc9562(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
