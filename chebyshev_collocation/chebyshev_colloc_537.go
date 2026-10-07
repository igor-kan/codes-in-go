package chebyshev_collocation

import "math"

// ComputeChebyshevColloc537 evaluates chebyshev collocation node order 537.
func ComputeChebyshevColloc537(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
