package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6032 evaluates chebyshev collocation node order 6032.
func ComputeChebyshevColloc6032(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
