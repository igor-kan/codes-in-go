package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8827 evaluates chebyshev collocation node order 8827.
func ComputeChebyshevColloc8827(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
