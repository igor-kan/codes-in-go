package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9732 evaluates chebyshev collocation node order 9732.
func ComputeChebyshevColloc9732(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
