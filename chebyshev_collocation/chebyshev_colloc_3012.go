package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3012 evaluates chebyshev collocation node order 3012.
func ComputeChebyshevColloc3012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
