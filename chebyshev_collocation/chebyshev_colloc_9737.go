package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9737 evaluates chebyshev collocation node order 9737.
func ComputeChebyshevColloc9737(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
