package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6052 evaluates chebyshev collocation node order 6052.
func ComputeChebyshevColloc6052(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
