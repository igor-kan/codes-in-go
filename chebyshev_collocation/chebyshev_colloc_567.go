package chebyshev_collocation

import "math"

// ComputeChebyshevColloc567 evaluates chebyshev collocation node order 567.
func ComputeChebyshevColloc567(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
