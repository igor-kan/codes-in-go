package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8532 evaluates chebyshev collocation node order 8532.
func ComputeChebyshevColloc8532(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
