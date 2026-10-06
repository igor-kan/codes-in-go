package chebyshev_collocation

import "math"

// ComputeChebyshevColloc112 evaluates chebyshev collocation node order 112.
func ComputeChebyshevColloc112(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
