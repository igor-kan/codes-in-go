package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8317 evaluates chebyshev collocation node order 8317.
func ComputeChebyshevColloc8317(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
