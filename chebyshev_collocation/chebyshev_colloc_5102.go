package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5102 evaluates chebyshev collocation node order 5102.
func ComputeChebyshevColloc5102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
