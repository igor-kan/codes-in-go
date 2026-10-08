package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1007 evaluates chebyshev collocation node order 1007.
func ComputeChebyshevColloc1007(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
