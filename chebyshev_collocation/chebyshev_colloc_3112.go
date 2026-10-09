package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3112 evaluates chebyshev collocation node order 3112.
func ComputeChebyshevColloc3112(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
