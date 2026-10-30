package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9287 evaluates chebyshev collocation node order 9287.
func ComputeChebyshevColloc9287(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
