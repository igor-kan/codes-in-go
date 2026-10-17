package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7747 evaluates chebyshev collocation node order 7747.
func ComputeChebyshevColloc7747(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
