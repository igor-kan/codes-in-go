package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8937 evaluates chebyshev collocation node order 8937.
func ComputeChebyshevColloc8937(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
