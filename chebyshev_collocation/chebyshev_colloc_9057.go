package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9057 evaluates chebyshev collocation node order 9057.
func ComputeChebyshevColloc9057(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
