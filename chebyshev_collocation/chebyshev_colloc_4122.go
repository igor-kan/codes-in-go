package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4122 evaluates chebyshev collocation node order 4122.
func ComputeChebyshevColloc4122(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
