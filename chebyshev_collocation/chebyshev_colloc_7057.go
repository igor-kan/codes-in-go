package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7057 evaluates chebyshev collocation node order 7057.
func ComputeChebyshevColloc7057(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
