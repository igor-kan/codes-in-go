package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9432 evaluates chebyshev collocation node order 9432.
func ComputeChebyshevColloc9432(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
