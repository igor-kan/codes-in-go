package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8152 evaluates chebyshev collocation node order 8152.
func ComputeChebyshevColloc8152(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
