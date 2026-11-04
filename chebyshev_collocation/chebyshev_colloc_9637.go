package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9637 evaluates chebyshev collocation node order 9637.
func ComputeChebyshevColloc9637(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
