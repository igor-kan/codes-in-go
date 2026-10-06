package chebyshev_collocation

import "math"

// ComputeChebyshevColloc32 evaluates chebyshev collocation node order 32.
func ComputeChebyshevColloc32(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
