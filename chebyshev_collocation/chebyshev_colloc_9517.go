package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9517 evaluates chebyshev collocation node order 9517.
func ComputeChebyshevColloc9517(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
