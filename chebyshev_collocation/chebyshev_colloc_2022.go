package chebyshev_collocation

import "math"

// ComputeChebyshevColloc2022 evaluates chebyshev collocation node order 2022.
func ComputeChebyshevColloc2022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
