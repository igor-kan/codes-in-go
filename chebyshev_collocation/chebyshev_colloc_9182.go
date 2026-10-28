package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9182 evaluates chebyshev collocation node order 9182.
func ComputeChebyshevColloc9182(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
