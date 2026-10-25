package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8907 evaluates chebyshev collocation node order 8907.
func ComputeChebyshevColloc8907(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
