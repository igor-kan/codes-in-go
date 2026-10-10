package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4002 evaluates chebyshev collocation node order 4002.
func ComputeChebyshevColloc4002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
