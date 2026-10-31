package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9352 evaluates chebyshev collocation node order 9352.
func ComputeChebyshevColloc9352(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
