package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5087 evaluates chebyshev collocation node order 5087.
func ComputeChebyshevColloc5087(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
