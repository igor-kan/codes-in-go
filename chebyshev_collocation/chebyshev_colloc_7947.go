package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7947 evaluates chebyshev collocation node order 7947.
func ComputeChebyshevColloc7947(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
