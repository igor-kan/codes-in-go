package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7422 evaluates chebyshev collocation node order 7422.
func ComputeChebyshevColloc7422(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
