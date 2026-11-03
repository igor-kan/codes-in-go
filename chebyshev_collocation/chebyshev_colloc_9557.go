package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9557 evaluates chebyshev collocation node order 9557.
func ComputeChebyshevColloc9557(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
