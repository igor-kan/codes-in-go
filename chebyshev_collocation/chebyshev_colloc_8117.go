package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8117 evaluates chebyshev collocation node order 8117.
func ComputeChebyshevColloc8117(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
