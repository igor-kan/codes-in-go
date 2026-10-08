package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1087 evaluates chebyshev collocation node order 1087.
func ComputeChebyshevColloc1087(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
