package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1122 evaluates chebyshev collocation node order 1122.
func ComputeChebyshevColloc1122(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
