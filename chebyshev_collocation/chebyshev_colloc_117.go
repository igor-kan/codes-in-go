package chebyshev_collocation

import "math"

// ComputeChebyshevColloc117 evaluates chebyshev collocation node order 117.
func ComputeChebyshevColloc117(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
