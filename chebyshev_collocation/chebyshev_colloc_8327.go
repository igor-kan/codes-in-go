package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8327 evaluates chebyshev collocation node order 8327.
func ComputeChebyshevColloc8327(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
