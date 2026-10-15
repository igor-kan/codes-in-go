package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7447 evaluates chebyshev collocation node order 7447.
func ComputeChebyshevColloc7447(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
