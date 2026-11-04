package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9692 evaluates chebyshev collocation node order 9692.
func ComputeChebyshevColloc9692(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
