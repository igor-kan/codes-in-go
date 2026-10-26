package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9002 evaluates chebyshev collocation node order 9002.
func ComputeChebyshevColloc9002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
