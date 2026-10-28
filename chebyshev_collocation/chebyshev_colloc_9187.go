package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9187 evaluates chebyshev collocation node order 9187.
func ComputeChebyshevColloc9187(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
