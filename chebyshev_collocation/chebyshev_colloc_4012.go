package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4012 evaluates chebyshev collocation node order 4012.
func ComputeChebyshevColloc4012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
