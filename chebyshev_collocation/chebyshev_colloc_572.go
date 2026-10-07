package chebyshev_collocation

import "math"

// ComputeChebyshevColloc572 evaluates chebyshev collocation node order 572.
func ComputeChebyshevColloc572(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
