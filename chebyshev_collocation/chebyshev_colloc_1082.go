package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1082 evaluates chebyshev collocation node order 1082.
func ComputeChebyshevColloc1082(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
