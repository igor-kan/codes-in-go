package chebyshev_collocation

import "math"

// ComputeChebyshevColloc502 evaluates chebyshev collocation node order 502.
func ComputeChebyshevColloc502(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
