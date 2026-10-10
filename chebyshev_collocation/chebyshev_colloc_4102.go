package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4102 evaluates chebyshev collocation node order 4102.
func ComputeChebyshevColloc4102(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
