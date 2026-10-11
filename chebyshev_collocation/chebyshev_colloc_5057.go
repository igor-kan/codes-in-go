package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5057 evaluates chebyshev collocation node order 5057.
func ComputeChebyshevColloc5057(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
