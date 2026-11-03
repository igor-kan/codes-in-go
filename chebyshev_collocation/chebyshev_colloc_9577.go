package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9577 evaluates chebyshev collocation node order 9577.
func ComputeChebyshevColloc9577(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
