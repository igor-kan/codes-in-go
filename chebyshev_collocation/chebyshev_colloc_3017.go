package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3017 evaluates chebyshev collocation node order 3017.
func ComputeChebyshevColloc3017(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
