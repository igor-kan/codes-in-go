package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8352 evaluates chebyshev collocation node order 8352.
func ComputeChebyshevColloc8352(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
