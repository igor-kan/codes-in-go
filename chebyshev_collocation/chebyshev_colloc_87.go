package chebyshev_collocation

import "math"

// ComputeChebyshevColloc87 evaluates chebyshev collocation node order 87.
func ComputeChebyshevColloc87(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
