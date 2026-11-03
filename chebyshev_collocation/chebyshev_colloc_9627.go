package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9627 evaluates chebyshev collocation node order 9627.
func ComputeChebyshevColloc9627(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
