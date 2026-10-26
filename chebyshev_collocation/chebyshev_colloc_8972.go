package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8972 evaluates chebyshev collocation node order 8972.
func ComputeChebyshevColloc8972(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
