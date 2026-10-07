package chebyshev_collocation

import "math"

// ComputeChebyshevColloc542 evaluates chebyshev collocation node order 542.
func ComputeChebyshevColloc542(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
