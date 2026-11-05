package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9767 evaluates chebyshev collocation node order 9767.
func ComputeChebyshevColloc9767(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
