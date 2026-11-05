package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9752 evaluates chebyshev collocation node order 9752.
func ComputeChebyshevColloc9752(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
