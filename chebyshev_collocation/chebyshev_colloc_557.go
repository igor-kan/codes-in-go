package chebyshev_collocation

import "math"

// ComputeChebyshevColloc557 evaluates chebyshev collocation node order 557.
func ComputeChebyshevColloc557(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
