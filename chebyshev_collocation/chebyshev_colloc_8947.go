package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8947 evaluates chebyshev collocation node order 8947.
func ComputeChebyshevColloc8947(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
