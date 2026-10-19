package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8112 evaluates chebyshev collocation node order 8112.
func ComputeChebyshevColloc8112(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
