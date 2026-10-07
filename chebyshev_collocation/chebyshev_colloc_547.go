package chebyshev_collocation

import "math"

// ComputeChebyshevColloc547 evaluates chebyshev collocation node order 547.
func ComputeChebyshevColloc547(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
