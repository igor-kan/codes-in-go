package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8157 evaluates chebyshev collocation node order 8157.
func ComputeChebyshevColloc8157(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
