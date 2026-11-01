package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9417 evaluates chebyshev collocation node order 9417.
func ComputeChebyshevColloc9417(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
