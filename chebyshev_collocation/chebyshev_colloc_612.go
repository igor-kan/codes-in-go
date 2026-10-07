package chebyshev_collocation

import "math"

// ComputeChebyshevColloc612 evaluates chebyshev collocation node order 612.
func ComputeChebyshevColloc612(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
