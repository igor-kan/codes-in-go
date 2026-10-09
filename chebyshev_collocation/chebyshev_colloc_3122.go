package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3122 evaluates chebyshev collocation node order 3122.
func ComputeChebyshevColloc3122(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
