package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8307 evaluates chebyshev collocation node order 8307.
func ComputeChebyshevColloc8307(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
