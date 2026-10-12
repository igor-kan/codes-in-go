package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6047 evaluates chebyshev collocation node order 6047.
func ComputeChebyshevColloc6047(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
