package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3117 evaluates chebyshev collocation node order 3117.
func ComputeChebyshevColloc3117(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
