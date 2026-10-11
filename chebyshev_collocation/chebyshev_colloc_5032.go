package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5032 evaluates chebyshev collocation node order 5032.
func ComputeChebyshevColloc5032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
