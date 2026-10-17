package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7737 evaluates chebyshev collocation node order 7737.
func ComputeChebyshevColloc7737(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
