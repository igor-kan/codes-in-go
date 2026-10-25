package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8962 evaluates chebyshev collocation node order 8962.
func ComputeChebyshevColloc8962(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
