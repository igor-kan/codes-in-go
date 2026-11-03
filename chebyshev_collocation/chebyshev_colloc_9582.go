package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9582 evaluates chebyshev collocation node order 9582.
func ComputeChebyshevColloc9582(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
