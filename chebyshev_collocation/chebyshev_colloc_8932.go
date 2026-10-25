package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8932 evaluates chebyshev collocation node order 8932.
func ComputeChebyshevColloc8932(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
