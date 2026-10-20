package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8337 evaluates chebyshev collocation node order 8337.
func ComputeChebyshevColloc8337(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
