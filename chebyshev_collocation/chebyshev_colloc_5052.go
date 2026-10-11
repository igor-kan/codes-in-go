package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5052 evaluates chebyshev collocation node order 5052.
func ComputeChebyshevColloc5052(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
