package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9107 evaluates chebyshev collocation node order 9107.
func ComputeChebyshevColloc9107(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
