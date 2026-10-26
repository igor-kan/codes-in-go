package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9022 evaluates chebyshev collocation node order 9022.
func ComputeChebyshevColloc9022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
