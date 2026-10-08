package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1057 evaluates chebyshev collocation node order 1057.
func ComputeChebyshevColloc1057(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
