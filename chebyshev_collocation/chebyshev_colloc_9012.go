package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9012 evaluates chebyshev collocation node order 9012.
func ComputeChebyshevColloc9012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
