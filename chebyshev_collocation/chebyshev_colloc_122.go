package chebyshev_collocation

import "math"

// ComputeChebyshevColloc122 evaluates chebyshev collocation node order 122.
func ComputeChebyshevColloc122(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
