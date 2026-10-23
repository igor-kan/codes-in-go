package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8717 evaluates chebyshev collocation node order 8717.
func ComputeChebyshevColloc8717(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
