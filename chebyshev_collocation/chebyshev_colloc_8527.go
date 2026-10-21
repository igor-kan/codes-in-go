package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8527 evaluates chebyshev collocation node order 8527.
func ComputeChebyshevColloc8527(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
