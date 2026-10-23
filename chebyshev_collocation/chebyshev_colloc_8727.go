package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8727 evaluates chebyshev collocation node order 8727.
func ComputeChebyshevColloc8727(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
