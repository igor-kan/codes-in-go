package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9787 evaluates chebyshev collocation node order 9787.
func ComputeChebyshevColloc9787(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
