package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1112 evaluates chebyshev collocation node order 1112.
func ComputeChebyshevColloc1112(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
