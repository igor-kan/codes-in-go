package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3002 evaluates chebyshev collocation node order 3002.
func ComputeChebyshevColloc3002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
