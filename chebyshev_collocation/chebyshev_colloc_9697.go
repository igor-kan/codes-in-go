package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9697 evaluates chebyshev collocation node order 9697.
func ComputeChebyshevColloc9697(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
