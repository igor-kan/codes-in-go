package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8367 evaluates chebyshev collocation node order 8367.
func ComputeChebyshevColloc8367(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
