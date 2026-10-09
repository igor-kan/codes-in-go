package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3027 evaluates chebyshev collocation node order 3027.
func ComputeChebyshevColloc3027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
