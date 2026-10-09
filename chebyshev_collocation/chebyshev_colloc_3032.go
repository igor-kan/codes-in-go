package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3032 evaluates chebyshev collocation node order 3032.
func ComputeChebyshevColloc3032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
