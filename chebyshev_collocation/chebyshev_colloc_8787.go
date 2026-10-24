package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8787 evaluates chebyshev collocation node order 8787.
func ComputeChebyshevColloc8787(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
