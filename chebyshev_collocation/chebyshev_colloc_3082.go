package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3082 evaluates chebyshev collocation node order 3082.
func ComputeChebyshevColloc3082(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
