package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5067 evaluates chebyshev collocation node order 5067.
func ComputeChebyshevColloc5067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
