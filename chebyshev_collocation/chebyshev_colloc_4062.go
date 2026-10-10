package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4062 evaluates chebyshev collocation node order 4062.
func ComputeChebyshevColloc4062(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
