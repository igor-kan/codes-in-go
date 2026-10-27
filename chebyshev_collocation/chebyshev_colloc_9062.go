package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9062 evaluates chebyshev collocation node order 9062.
func ComputeChebyshevColloc9062(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
