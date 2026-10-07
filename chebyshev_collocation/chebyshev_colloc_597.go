package chebyshev_collocation

import "math"

// ComputeChebyshevColloc597 evaluates chebyshev collocation node order 597.
func ComputeChebyshevColloc597(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
