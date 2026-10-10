package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4087 evaluates chebyshev collocation node order 4087.
func ComputeChebyshevColloc4087(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
