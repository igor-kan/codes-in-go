package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1022 evaluates chebyshev collocation node order 1022.
func ComputeChebyshevColloc1022(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
