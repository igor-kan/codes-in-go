package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7247 evaluates chebyshev collocation node order 7247.
func ComputeChebyshevColloc7247(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
