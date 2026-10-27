package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9147 evaluates chebyshev collocation node order 9147.
func ComputeChebyshevColloc9147(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
