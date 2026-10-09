package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3042 evaluates chebyshev collocation node order 3042.
func ComputeChebyshevColloc3042(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
