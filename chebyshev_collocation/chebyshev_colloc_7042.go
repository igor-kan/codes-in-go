package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7042 evaluates chebyshev collocation node order 7042.
func ComputeChebyshevColloc7042(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
