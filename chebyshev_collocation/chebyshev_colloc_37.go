package chebyshev_collocation

import "math"

// ComputeChebyshevColloc37 evaluates chebyshev collocation node order 37.
func ComputeChebyshevColloc37(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
