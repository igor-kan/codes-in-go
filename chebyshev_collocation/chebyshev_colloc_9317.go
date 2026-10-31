package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9317 evaluates chebyshev collocation node order 9317.
func ComputeChebyshevColloc9317(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
