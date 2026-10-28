package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9217 evaluates chebyshev collocation node order 9217.
func ComputeChebyshevColloc9217(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
